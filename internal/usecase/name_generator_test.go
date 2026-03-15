package usecase_test

import (
	"math/rand"
	"testing"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
	"github.com/code-rcplaza/rpg_engine/internal/usecase"
)

// mockNameRepo implements NameRepository with hardcoded data.
// No DB, no network — tests run in microseconds.
// Go needs no mocking libraries: if it satisfies the interface, it works.
type mockNameRepo struct {
	races      map[string]domain.Race
	patterns   map[int][]domain.NamePattern
	components map[int]map[string][]domain.NameComponent
	composites map[int]map[string][]domain.CompositePart
}

func (m *mockNameRepo) FindRace(slug string) (domain.Race, error) {
	race, ok := m.races[slug]
	if !ok {
		return domain.Race{}, &domain.NotFoundError{Entity: "race", ID: slug}
	}
	return race, nil
}

func (m *mockNameRepo) FindPatterns(raceID int) ([]domain.NamePattern, error) {
	return m.patterns[raceID], nil
}

func (m *mockNameRepo) FindComponents(raceID int, componentType string, gender domain.Gender) ([]domain.NameComponent, error) {
	byType, ok := m.components[raceID]
	if !ok {
		return nil, nil
	}
	return byType[componentType], nil
}

func (m *mockNameRepo) FindCompositeParts(raceID int, position string) ([]domain.CompositePart, error) {
	byPos, ok := m.composites[raceID]
	if !ok {
		return nil, nil
	}
	return byPos[position], nil
}

// newSeededGenerator creates a generator with a fixed seed.
// Same seed → same result → deterministic tests.
func newSeededGenerator(repo usecase.NameRepository) *usecase.NameGenerator {
	return usecase.NewNameGenerator(repo, rand.New(rand.NewSource(42)))
}

// --- Tests ---

// TestGenerate uses table-driven tests: the idiomatic Go pattern.
// A single loop covers all cases — easy to extend.
func TestGenerate(t *testing.T) {
	repo := buildTestRepo()

	tests := []struct {
		name      string
		raceSlug  string
		gender    domain.Gender
		wantParts int
		wantErr   bool
	}{
		{
			name:      "human male generates first name and last name",
			raceSlug:  "human",
			gender:    domain.GenderMale,
			wantParts: 2,
			wantErr:   false,
		},
		{
			name:      "halfling generates first name and composite nickname",
			raceSlug:  "halfling",
			gender:    domain.GenderNeutral,
			wantParts: 2,
			wantErr:   false,
		},
		{
			name:     "unknown race returns error",
			raceSlug: "cosmic_dragon",
			gender:   domain.GenderNeutral,
			wantErr:  true,
		},
		{
			name:     "empty slug returns error",
			raceSlug: "",
			gender:   domain.GenderNeutral,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := newSeededGenerator(repo)

			result, err := gen.Generate(tt.raceSlug, tt.gender)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantErr {
				return
			}

			if len(result.Parts) != tt.wantParts {
				t.Errorf("expected %d parts, got %d: %v", tt.wantParts, len(result.Parts), result.Parts)
			}

			if result.Full == "" {
				t.Error("full name cannot be empty")
			}
		})
	}
}

func TestGenerate_FullNameJoinsParts(t *testing.T) {
	repo := buildTestRepo()
	gen := newSeededGenerator(repo)

	result, err := gen.Generate("human", domain.GenderMale)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := result.Parts[0] + " " + result.Parts[1]
	if result.Full != expected {
		t.Errorf("Full=%q does not match joined Parts=%q", result.Full, expected)
	}
}

// buildTestRepo builds the mock with test data.
// Kept separate so multiple tests can reuse it.
func buildTestRepo() *mockNameRepo {
	return &mockNameRepo{
		races: map[string]domain.Race{
			"human":    {ID: 1, Slug: "human", Name: "Human"},
			"halfling": {ID: 2, Slug: "halfling", Name: "Halfling"},
		},
		patterns: map[int][]domain.NamePattern{
			1: { // human
				{ID: 1, RaceID: 1, Order: 1, ComponentType: "first_name", Required: true},
				{ID: 2, RaceID: 1, Order: 2, ComponentType: "last_name", Required: true},
			},
			2: { // halfling
				{ID: 3, RaceID: 2, Order: 1, ComponentType: "first_name", Required: true},
				{ID: 4, RaceID: 2, Order: 2, ComponentType: "composite_nickname", Required: true},
			},
		},
		components: map[int]map[string][]domain.NameComponent{
			1: { // human
				"first_name": {
					{ID: 1, Value: "John"},
					{ID: 2, Value: "Peter"},
					{ID: 3, Value: "Diego"},
				},
				"last_name": {
					{ID: 4, Value: "Johnson"},
					{ID: 5, Value: "Smith"},
					{ID: 6, Value: "Martinez"},
				},
			},
			2: { // halfling
				"first_name": {
					{ID: 7, Value: "Fosco"},
					{ID: 8, Value: "Bilbo"},
					{ID: 9, Value: "Drogo"},
				},
			},
		},
		composites: map[int]map[string][]domain.CompositePart{
			2: { // halfling
				"first": {
					{ID: 1, Value: "Hot"},
					{ID: 2, Value: "Barrel"},
					{ID: 3, Value: "Pipe"},
				},
				"second": {
					{ID: 4, Value: "Pot"},
					{ID: 5, Value: "Old"},
					{ID: 6, Value: "Broken"},
				},
			},
		},
	}
}