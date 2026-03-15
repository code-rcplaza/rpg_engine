package usecase_test

import (
	"math/rand"
	"testing"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
	"github.com/code-rcplaza/rpg_engine/internal/usecase"
)

// mockNameRepo implements NameRepository with hardcoded data.
// No DB, no network — tests run in microseconds.
type mockNameRepo struct {
	races      map[string]domain.Race
	styles     map[int][]domain.NameStyle
	patterns   map[int][]domain.NamePattern
	byStyle    map[int][]domain.NamePattern
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

func (m *mockNameRepo) FindStyles(raceID int) ([]domain.NameStyle, error) {
	return m.styles[raceID], nil
}

func (m *mockNameRepo) FindPatterns(raceID int) ([]domain.NamePattern, error) {
	return m.patterns[raceID], nil
}

func (m *mockNameRepo) FindPatternsByStyle(raceID int, styleID int) ([]domain.NamePattern, error) {
	return m.byStyle[styleID], nil
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
		},
		{
			name:      "halfling generates first name and composite last name",
			raceSlug:  "halfling",
			gender:    domain.GenderNeutral,
			wantParts: 2,
		},
		{
			name:      "tiefling picks one style and generates one name part",
			raceSlug:  "tiefling",
			gender:    domain.GenderMale,
			wantParts: 1,
		},
		{
			name:    "unknown race returns error",
			raceSlug: "cosmic_dragon",
			gender:  domain.GenderNeutral,
			wantErr: true,
		},
		{
			name:    "empty slug returns error",
			raceSlug: "",
			gender:  domain.GenderNeutral,
			wantErr: true,
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

func TestGenerate_HalflingCompositeLastName(t *testing.T) {
	repo := buildTestRepo()
	gen := newSeededGenerator(repo)

	result, err := gen.Generate("halfling", domain.GenderMale)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Parts) < 2 {
		t.Fatalf("expected at least 2 parts, got %d", len(result.Parts))
	}

	// Last name must not contain a space (it is a compound word, not two words)
	lastName := result.Parts[1]
	for _, c := range lastName {
		if c == ' ' {
			t.Errorf("halfling last name should be one compound word, got: %q", lastName)
		}
	}
}

// buildTestRepo builds the mock with test data covering all cases.
func buildTestRepo() *mockNameRepo {
	tieflingInfernalStyleID := 1
	tieflingVirtueStyleID := 2

	return &mockNameRepo{
		races: map[string]domain.Race{
			"human":    {ID: 1, Slug: "human", Name: "Human"},
			"halfling": {ID: 2, Slug: "halfling", Name: "Halfling"},
			"tiefling": {ID: 3, Slug: "tiefling", Name: "Tiefling"},
		},
		styles: map[int][]domain.NameStyle{
			3: { // tiefling has two styles
				{ID: tieflingInfernalStyleID, RaceID: 3, Slug: "infernal"},
				{ID: tieflingVirtueStyleID, RaceID: 3, Slug: "virtue"},
			},
		},
		patterns: map[int][]domain.NamePattern{
			1: { // human — no styles
				{ID: 1, RaceID: 1, Order: 1, ComponentType: "first_name", Required: true, MaxCount: 1},
				{ID: 2, RaceID: 1, Order: 2, ComponentType: "last_name", Required: true, MaxCount: 1},
			},
			2: { // halfling — no styles
				{ID: 3, RaceID: 2, Order: 1, ComponentType: "first_name", Required: true, MaxCount: 1},
				{ID: 4, RaceID: 2, Order: 2, ComponentType: "last_name", Required: true, MaxCount: 1},
			},
		},
		byStyle: map[int][]domain.NamePattern{
			tieflingInfernalStyleID: {
				{ID: 5, RaceID: 3, StyleID: &tieflingInfernalStyleID, Order: 1, ComponentType: "infernal_name", Required: true, MaxCount: 1},
			},
			tieflingVirtueStyleID: {
				{ID: 6, RaceID: 3, StyleID: &tieflingVirtueStyleID, Order: 1, ComponentType: "virtue_name", Required: true, MaxCount: 1},
			},
		},
		components: map[int]map[string][]domain.NameComponent{
			1: {
				"first_name": {{ID: 1, Value: "John"}, {ID: 2, Value: "Peter"}},
				"last_name":  {{ID: 3, Value: "Smith"}, {ID: 4, Value: "Johnson"}},
			},
			2: {
				"first_name": {{ID: 5, Value: "Fosco"}, {ID: 6, Value: "Bilbo"}},
			},
			3: {
				"infernal_name": {{ID: 7, Value: "Akmenos"}, {ID: 8, Value: "Amnon"}},
				"virtue_name":   {{ID: 9, Value: "Hope"}, {ID: 10, Value: "Despair"}},
			},
		},
		composites: map[int]map[string][]domain.CompositePart{
			2: { // halfling
				"first":  {{ID: 1, Value: "Brush"}, {ID: 2, Value: "Good"}},
				"second": {{ID: 3, Value: "gather"}, {ID: 4, Value: "barrel"}},
			},
		},
	}
}