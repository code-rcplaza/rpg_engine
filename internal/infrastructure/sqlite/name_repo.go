package sqlite

import (
	"database/sql"
	"errors"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
	"github.com/code-rcplaza/rpg_engine/internal/usecase"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3" // driver registration side effect
)

// Compile-time check: NameRepo must satisfy usecase.NameRepository.
// If it doesn't, the build fails here with a clear error — not at runtime.
var _ usecase.NameRepository = (*NameRepo)(nil)

// NameRepo is the SQLite implementation of usecase.NameRepository.
type NameRepo struct {
	db *sqlx.DB
}

// NewNameRepo opens the SQLite database and returns a ready NameRepo.
func NewNameRepo(dbPath string) (*NameRepo, error) {
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &NameRepo{db: db}, nil
}

// Close closes the underlying database connection.
func (r *NameRepo) Close() error {
	return r.db.Close()
}

// FindRace retrieves a race by its slug.
func (r *NameRepo) FindRace(slug string) (domain.Race, error) {
	var row struct {
		ID   int    `db:"id"`
		Slug string `db:"slug"`
		Name string `db:"name"`
	}

	err := r.db.Get(&row, `SELECT id, slug, name FROM races WHERE slug = ?`, slug)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Race{}, &domain.NotFoundError{Entity: "race", ID: slug}
	}
	if err != nil {
		return domain.Race{}, err
	}

	return domain.Race{ID: row.ID, Slug: row.Slug, Name: row.Name}, nil
}

// FindStyles retrieves all name styles for a race.
// Returns empty slice if the race has no styles (most races).
func (r *NameRepo) FindStyles(raceID int) ([]domain.NameStyle, error) {
	var rows []struct {
		ID     int    `db:"id"`
		RaceID int    `db:"race_id"`
		Slug   string `db:"slug"`
	}

	err := r.db.Select(&rows, `
		SELECT id, race_id, slug
		FROM name_styles
		WHERE race_id = ?
	`, raceID)
	if err != nil {
		return nil, err
	}

	styles := make([]domain.NameStyle, len(rows))
	for i, row := range rows {
		styles[i] = domain.NameStyle{ID: row.ID, RaceID: row.RaceID, Slug: row.Slug}
	}

	return styles, nil
}

// FindPatterns retrieves all patterns for a race that have no style (style_id IS NULL).
func (r *NameRepo) FindPatterns(raceID int) ([]domain.NamePattern, error) {
	return r.queryPatterns(`
		SELECT id, race_id, style_id, "order", component_type, required, max_count
		FROM name_patterns
		WHERE race_id = ? AND style_id IS NULL
		ORDER BY "order" ASC
	`, raceID)
}

// FindPatternsByStyle retrieves patterns for a specific style.
func (r *NameRepo) FindPatternsByStyle(raceID int, styleID int) ([]domain.NamePattern, error) {
	return r.queryPatterns(`
		SELECT id, race_id, style_id, "order", component_type, required, max_count
		FROM name_patterns
		WHERE race_id = ? AND style_id = ?
		ORDER BY "order" ASC
	`, raceID, styleID)
}

// queryPatterns is a private helper that runs a pattern query and maps the rows.
func (r *NameRepo) queryPatterns(query string, args ...any) ([]domain.NamePattern, error) {
	var rows []struct {
		ID            int     `db:"id"`
		RaceID        int     `db:"race_id"`
		StyleID       *int    `db:"style_id"` // pointer because it's nullable
		Order         int     `db:"order"`
		ComponentType string  `db:"component_type"`
		Required      int     `db:"required"`
		MaxCount      int     `db:"max_count"`
	}

	err := r.db.Select(&rows, query, args...)
	if err != nil {
		return nil, err
	}

	patterns := make([]domain.NamePattern, len(rows))
	for i, row := range rows {
		patterns[i] = domain.NamePattern{
			ID:            row.ID,
			RaceID:        row.RaceID,
			StyleID:       row.StyleID,
			Order:         row.Order,
			ComponentType: row.ComponentType,
			Required:      row.Required == 1,
			MaxCount:      row.MaxCount,
		}
	}

	return patterns, nil
}

// FindComponents retrieves name components filtered by race, type and gender.
func (r *NameRepo) FindComponents(raceID int, componentType string, gender domain.Gender) ([]domain.NameComponent, error) {
	var rows []struct {
		ID            int    `db:"id"`
		RaceID        int    `db:"race_id"`
		ComponentType string `db:"component_type"`
		Gender        string `db:"gender"`
		Value         string `db:"value"`
	}

	var err error
	if gender == domain.GenderNeutral {
		err = r.db.Select(&rows, `
			SELECT id, race_id, component_type, COALESCE(gender, 'n') as gender, value
			FROM name_components
			WHERE race_id = ? AND component_type = ?
		`, raceID, componentType)
	} else {
		err = r.db.Select(&rows, `
			SELECT id, race_id, component_type, COALESCE(gender, 'n') as gender, value
			FROM name_components
			WHERE race_id = ? AND component_type = ? AND (gender = ? OR gender IS NULL)
		`, raceID, componentType, string(gender))
	}

	if err != nil {
		return nil, err
	}

	components := make([]domain.NameComponent, len(rows))
	for i, row := range rows {
		components[i] = domain.NameComponent{
			ID:            row.ID,
			RaceID:        row.RaceID,
			ComponentType: row.ComponentType,
			Gender:        domain.Gender(row.Gender),
			Value:         row.Value,
		}
	}

	return components, nil
}

// FindCompositeParts retrieves composite parts for a race by position.
func (r *NameRepo) FindCompositeParts(raceID int, position string) ([]domain.CompositePart, error) {
	var rows []struct {
		ID       int    `db:"id"`
		RaceID   int    `db:"race_id"`
		Position string `db:"position"`
		Category string `db:"category"`
		Value    string `db:"value"`
	}

	err := r.db.Select(&rows, `
		SELECT id, race_id, position, category, value
		FROM composite_parts
		WHERE race_id = ? AND position = ?
	`, raceID, position)
	if err != nil {
		return nil, err
	}

	parts := make([]domain.CompositePart, len(rows))
	for i, row := range rows {
		parts[i] = domain.CompositePart{
			ID:       row.ID,
			RaceID:   row.RaceID,
			Position: row.Position,
			Category: row.Category,
			Value:    row.Value,
		}
	}

	return parts, nil
}