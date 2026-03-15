package usecase

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
)

// NameRepository defines what the usecase needs from the data layer.
// The interface lives here (in usecase), not in infrastructure.
// That is dependency inversion: usecase does not know about SQLite.
type NameRepository interface {
	FindRace(slug string) (domain.Race, error)
	FindPatterns(raceID int) ([]domain.NamePattern, error)
	FindComponents(raceID int, componentType string, gender domain.Gender) ([]domain.NameComponent, error)
	FindCompositeParts(raceID int, position string) ([]domain.CompositePart, error)
}

// NameGenerator holds the name generation logic.
// It depends on the interface, never on a concrete implementation.
type NameGenerator struct {
	repo NameRepository
	rand *rand.Rand // injected so tests can be deterministic
}

// NewNameGenerator is the constructor. Go constructors are New* functions.
func NewNameGenerator(repo NameRepository, r *rand.Rand) *NameGenerator {
	return &NameGenerator{repo: repo, rand: r}
}

// Generate produces a full name for the given race slug and gender.
// Notice the early returns: each error exits immediately.
func (g *NameGenerator) Generate(raceSlug string, gender domain.Gender) (domain.GeneratedName, error) {
	if raceSlug == "" {
		return domain.GeneratedName{}, errors.New("race slug cannot be empty")
	}

	race, err := g.repo.FindRace(raceSlug)
	if err != nil {
		return domain.GeneratedName{}, err // early return
	}

	patterns, err := g.repo.FindPatterns(race.ID)
	if err != nil {
		return domain.GeneratedName{}, err // early return
	}

	if len(patterns) == 0 {
		return domain.GeneratedName{}, errors.New("no name pattern defined for this race")
	}

	parts, err := g.buildParts(race, patterns, gender)
	if err != nil {
		return domain.GeneratedName{}, err // early return
	}

	return domain.GeneratedName{
		Full:  strings.Join(parts, " "),
		Parts: parts,
		Race:  race,
	}, nil
}

// buildParts builds each part of the name according to the race pattern.
// Private function (lowercase) — only used within this package.
func (g *NameGenerator) buildParts(
	race domain.Race,
	patterns []domain.NamePattern,
	gender domain.Gender,
) ([]string, error) {
	var parts []string // zero value of slice is nil, append works fine

	for _, pattern := range patterns {
		// Skip non-required components randomly
		if !pattern.Required && g.rand.Intn(2) == 0 {
			continue
		}

		part, err := g.buildComponent(race.ID, pattern.ComponentType, gender)
		if err != nil {
			return nil, err // early return inside loop
		}

		parts = append(parts, part)
	}

	return parts, nil
}

// buildComponent resolves a single name component.
// Delegates to buildComposite for "composite_nickname", otherwise queries components.
func (g *NameGenerator) buildComponent(
	raceID int,
	componentType string,
	gender domain.Gender,
) (string, error) {
	if componentType == "composite_nickname" {
		return g.buildComposite(raceID)
	}

	candidates, err := g.repo.FindComponents(raceID, componentType, gender)
	if err != nil {
		return "", err
	}

	if len(candidates) == 0 {
		return "", errors.New("no components found for: " + componentType)
	}

	return pickRandom(g.rand, candidates).Value, nil
}

// buildComposite builds the compound nickname for races like Halflings.
// "first" + "second" → "Hot" + "Pot" → "Hot Pot"
func (g *NameGenerator) buildComposite(raceID int) (string, error) {
	first, err := g.repo.FindCompositeParts(raceID, "first")
	if err != nil {
		return "", err
	}

	second, err := g.repo.FindCompositeParts(raceID, "second")
	if err != nil {
		return "", err
	}

	if len(first) == 0 || len(second) == 0 {
		return "", errors.New("not enough composite parts for this race")
	}

	return pickRandom(g.rand, first).Value + " " + pickRandom(g.rand, second).Value, nil
}

// pickRandom selects a random element from a slice.
// Standalone function (not a method) because Go does not allow
// type parameters on struct methods.
func pickRandom[T any](r *rand.Rand, items []T) T {
	return items[r.Intn(len(items))]
}