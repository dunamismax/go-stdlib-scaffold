package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dunamismax/go-stdlib-scaffold/internal/database"
)

func main() {
	// Define command-line flags
	dbPath := flag.String("db", "data/app.db", "Path to the SQLite database file.")
	migrationsDir := flag.String("migrations", "internal/database/migrations", "Directory containing SQL migration files.")
	flag.Parse()

	// Get the command
	if flag.NArg() == 0 {
		log.Fatal("Usage: cli <command>\nCommands: migrate")
	}
	command := flag.Arg(0)

	// Execute the command
	switch command {
	case "migrate":
		if err := runMigrations(*dbPath, *migrationsDir); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migrations applied successfully.")
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

func runMigrations(dbPath, migrationsDir string) error {
	db, err := database.NewDB(dbPath)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	if err := database.Migrate(db, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}