package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/code-rcplaza/rpg_engine/internal/infrastructure/sqlite"
	"github.com/code-rcplaza/rpg_engine/internal/usecase"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/rpggen.db" // fallback for local dev outside Docker
	}

	// Infrastructure layer — the only place that knows about SQLite
	repo, err := sqlite.NewNameRepo(dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer repo.Close()

	// Usecase layer — receives the repo through the interface
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	generator := usecase.NewNameGenerator(repo, rng)

	// Quick smoke test — remove once GraphQL is wired up
	result, err := generator.Generate("human", "m")
	if err != nil {
		log.Fatalf("generation failed: %v", err)
	}

	log.Printf("generated name: %s", result.Full)
}