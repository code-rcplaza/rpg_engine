package usecase_test

import (
	"math/rand"
	"testing"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
	"github.comcode-rcplaza/rpg_engine/internal/usecase"
)

// mockNameRepo implementa NameRepository con datos hardcodeados.
// No hay DB, no hay red — los tests corren en microsegundos.
// Go no necesita librerías de mocking: si satisface la interfaz, funciona.
type mockNameRepo struct {
	races      map[string]domain.Race
	patterns   map[int][]domain.NamePattern
	components map[int]map[string][]domain.NameComponent
	composites map[int]map[string][]domain.CompositePart
}

func (m *mockNameRepo) FindRace(slug string) (domain.Race, error) {
	race, ok := m.races[slug]
	if !ok {
		return domain.Race{}, &domain.NotFoundError{Entity: "raza", ID: slug}
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

// newSeededGenerator crea un generador con semilla fija.
// Mismo seed → mismo resultado → tests deterministas.
func newSeededGenerator(repo usecase.NameRepository) *usecase.NameGenerator {
	return usecase.NewNameGenerator(repo, rand.New(rand.NewSource(42)))
}

// --- Tests ---

// TestGenerate usa table-driven tests: el patrón idiomático de Go.
// Un solo loop cubre todos los casos — fácil de extender.
func TestGenerate(t *testing.T) {
	repo := buildTestRepo()

	tests := []struct {
		name      string // descripción del caso
		raceSlug  string
		gender    domain.Gender
		wantParts int  // cuántas partes esperamos en el nombre
		wantErr   bool
	}{
		{
			name:      "humano masculino genera nombre y apellido",
			raceSlug:  "humano",
			gender:    domain.GenderMale,
			wantParts: 2,
			wantErr:   false,
		},
		{
			name:      "mediano genera nombre y apodo compuesto",
			raceSlug:  "mediano",
			gender:    domain.GenderNeutral,
			wantParts: 2,
			wantErr:   false,
		},
		{
			name:     "raza inexistente retorna error",
			raceSlug: "dragon_cosmico",
			gender:   domain.GenderNeutral,
			wantErr:  true,
		},
		{
			name:     "slug vacío retorna error",
			raceSlug: "",
			gender:   domain.GenderNeutral,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := newSeededGenerator(repo)

			result, err := gen.Generate(tt.raceSlug, tt.gender)

			// Verificar error
			if tt.wantErr && err == nil {
				t.Fatal("esperaba error pero no hubo ninguno")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("no esperaba error, pero obtuvo: %v", err)
			}

			// Si se esperaba error no verificamos el resultado
			if tt.wantErr {
				return
			}

			// Verificar partes del nombre
			if len(result.Parts) != tt.wantParts {
				t.Errorf("esperaba %d partes, obtuvo %d: %v", tt.wantParts, len(result.Parts), result.Parts)
			}

			// El nombre completo nunca puede estar vacío si no hubo error
			if result.Full == "" {
				t.Error("nombre completo no puede estar vacío")
			}
		})
	}
}

func TestGenerate_NombreCompletoUnePartes(t *testing.T) {
	repo := buildTestRepo()
	gen := newSeededGenerator(repo)

	result, err := gen.Generate("humano", domain.GenderMale)
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}

	expected := result.Parts[0] + " " + result.Parts[1]
	if result.Full != expected {
		t.Errorf("Full=%q no coincide con Parts unidos=%q", result.Full, expected)
	}
}

// buildTestRepo construye el mock con datos de prueba.
// Separado para reusar en múltiples tests.
func buildTestRepo() *mockNameRepo {
	return &mockNameRepo{
		races: map[string]domain.Race{
			"humano":  {ID: 1, Slug: "humano", Name: "Humano"},
			"mediano": {ID: 2, Slug: "mediano", Name: "Mediano"},
		},
		patterns: map[int][]domain.NamePattern{
			1: { // humano
				{ID: 1, RaceID: 1, Order: 1, ComponentType: "nombre_pila", Required: true},
				{ID: 2, RaceID: 1, Order: 2, ComponentType: "apellido", Required: true},
			},
			2: { // mediano
				{ID: 3, RaceID: 2, Order: 1, ComponentType: "nombre_pila", Required: true},
				{ID: 4, RaceID: 2, Order: 2, ComponentType: "apodo_compuesto", Required: true},
			},
		},
		components: map[int]map[string][]domain.NameComponent{
			1: { // humano
				"nombre_pila": {
					{ID: 1, Value: "Juan"},
					{ID: 2, Value: "Pedro"},
					{ID: 3, Value: "Diego"},
				},
				"apellido": {
					{ID: 4, Value: "García"},
					{ID: 5, Value: "López"},
					{ID: 6, Value: "Martínez"},
				},
			},
			2: { // mediano
				"nombre_pila": {
					{ID: 7, Value: "Fosco"},
					{ID: 8, Value: "Bilbo"},
					{ID: 9, Value: "Drogo"},
				},
			},
		},
		composites: map[int]map[string][]domain.CompositePart{
			2: { // mediano
				"primera": {
					{ID: 1, Value: "Olla"},
					{ID: 2, Value: "Barril"},
					{ID: 3, Value: "Pipa"},
				},
				"segunda": {
					{ID: 4, Value: "Caliente"},
					{ID: 5, Value: "Viejo"},
					{ID: 6, Value: "Roto"},
				},
			},
		},
	}
}