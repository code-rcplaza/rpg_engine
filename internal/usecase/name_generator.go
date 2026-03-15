package usecase

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
)

// NameRepository define QUÉ necesita el usecase de la capa de datos.
// La interfaz vive acá (en usecase), no en infrastructure.
// Eso es la inversión de dependencias: usecase no conoce SQLite.
type NameRepository interface {
	FindRace(slug string) (domain.Race, error)
	FindPatterns(raceID int) ([]domain.NamePattern, error)
	FindComponents(raceID int, componentType string, gender domain.Gender) ([]domain.NameComponent, error)
	FindCompositeParts(raceID int, position string) ([]domain.CompositePart, error)
}

// NameGenerator contiene la lógica de generación.
// Depende de la interfaz, nunca de una implementación concreta.
type NameGenerator struct {
	repo NameRepository
	rand *rand.Rand // inyectado para que los tests sean deterministas
}

// NewNameGenerator es el constructor. En Go los constructores son funciones New*.
func NewNameGenerator(repo NameRepository, r *rand.Rand) *NameGenerator {
	return &NameGenerator{repo: repo, rand: r}
}

// Generate genera un nombre completo para la raza y género dados.
// Nota los early returns: cada error sale inmediatamente.
func (g *NameGenerator) Generate(raceSlug string, gender domain.Gender) (domain.GeneratedName, error) {
	if raceSlug == "" {
		return domain.GeneratedName{}, errors.New("raza no puede estar vacía")
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
		return domain.GeneratedName{}, errors.New("raza sin patrón de nombre definido")
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

// buildParts construye cada parte del nombre según el patrón de la raza.
// Función privada (minúscula) — solo la usa este paquete.
func (g *NameGenerator) buildParts(
	race domain.Race,
	patterns []domain.NamePattern,
	gender domain.Gender,
) ([]string, error) {
	var parts []string // zero value de slice es nil, append funciona igual

	for _, pattern := range patterns {
		// Si el componente no es obligatorio, lo saltamos aleatoriamente
		if !pattern.Required && g.rand.Intn(2) == 0 {
			continue
		}

		part, err := g.buildComponent(race.ID, pattern.ComponentType, gender)
		if err != nil {
			return nil, err // early return dentro del loop
		}

		parts = append(parts, part)
	}

	return parts, nil
}

// buildComponent resuelve un único componente del nombre.
// Si es "apodo_compuesto" delega a buildComposite, si no busca en componentes.
func (g *NameGenerator) buildComponent(
	raceID int,
	componentType string,
	gender domain.Gender,
) (string, error) {
	if componentType == "apodo_compuesto" {
		return g.buildComposite(raceID)
	}

	// Para componentes normales buscamos candidatos y elegimos uno al azar
	candidates, err := g.repo.FindComponents(raceID, componentType, gender)
	if err != nil {
		return "", err
	}

	if len(candidates) == 0 {
		return "", errors.New("sin componentes para: " + componentType)
	}

	return g.pickRandom(candidates).Value, nil
}

// buildComposite arma el apodo compuesto de razas como el Mediano.
// "primera" + "segunda" → "Olla" + "Caliente" → "Olla Caliente"
func (g *NameGenerator) buildComposite(raceID int) (string, error) {
	primera, err := g.repo.FindCompositeParts(raceID, "primera")
	if err != nil {
		return "", err
	}

	segunda, err := g.repo.FindCompositeParts(raceID, "segunda")
	if err != nil {
		return "", err
	}

	if len(primera) == 0 || len(segunda) == 0 {
		return "", errors.New("partes compuestas insuficientes para esta raza")
	}

	return g.pickRandom(primera).Value + " " + g.pickRandom(segunda).Value, nil
}

// pickRandom elige un elemento aleatorio de un slice.
// Función genérica — el [T any] significa que acepta cualquier tipo.
// Esto es programación funcional en Go: función pura, sin side effects.
func (g *NameGenerator) pickRandom[T any](items []T) T {
	return items[g.rand.Intn(len(items))]
}