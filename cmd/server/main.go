package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/code-rcplaza/rpg_engine/internal/domain"
	"github.com/code-rcplaza/rpg_engine/internal/infrastructure/sqlite"
	"github.com/code-rcplaza/rpg_engine/internal/usecase"
)

type testCase struct {
	race   string
	gender domain.Gender
}

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/rpggen.db"
	}

	repo, err := sqlite.NewNameRepo(dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer repo.Close()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	generator := usecase.NewNameGenerator(repo, rng)

	tests := []testCase{
		{"human", "m"},
		{"human", "f"},
		{"halfling", "m"},
		{"halfling", "f"},
		{"half-elf", "m"},
		{"half-elf", "f"},
		{"half-orc", "m"},
		{"half-orc", "f"},
		{"dragonborn", "m"},
		{"dragonborn", "f"},
		{"gnome", "m"},
		{"gnome", "f"},
		{"tiefling", "m"},
		{"tiefling", "f"},
	}

	for _, tc := range tests {
		result, err := generator.Generate(tc.race, tc.gender)
		if err != nil {
			log.Fatalf("generation failed for race=%s gender=%s: %v", tc.race, tc.gender, err)
		}

		log.Printf("race=%s gender=%s -> %s", tc.race, tc.gender, result.Full)
	}
}