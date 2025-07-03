package cli

import (
	"fmt"
	"os"
	"strings"

	cmdServer "basic-crud-go/cmd/server"
	"basic-crud-go/internal/infrastructure/db/postgres"
)

func HandleCLI() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := strings.TrimSpace(os.Args[1])

	switch command {
	case "--start":
		fmt.Println("🚀 Starting application...")
		// Inicialize connection database
		postgres.InitPostgres()
		// Inicialize server
		router := cmdServer.InitServer()
		cmdServer.StartServer(router)

	case "--db-check":
		fmt.Println("⏳ Checking database connection...")
		postgres.InitPostgres()

	case "--db-create":
		fmt.Println("🔨 Creating database structure...")
		postgres.InitPostgres()
		postgres.RunMigrations()

	case "--db-drop":
		if len(os.Args) < 3 {
			fmt.Println("❗ Please provide the folder name to rollback. Example: --db-drop 03_middleware")
			return
		}
		folder := os.Args[2]
		fmt.Printf("🧨 Dropping tables from folder: %s...\n", folder)
		postgres.InitPostgres()
		postgres.RollbackByFolder(folder)

	case "--status":
		db := postgres.InitPostgres()
		rows, err := db.Query(`SELECT filename, applied_at FROM migrations_applied ORDER BY applied_at`)
		if err != nil {
			fmt.Printf("❌ Failed to query migrations: %v\n", err)
			return
		}
		defer rows.Close()

		fmt.Println("📄 Applied migrations:")
		for rows.Next() {
			var name string
			var appliedAt string
			if err := rows.Scan(&name, &appliedAt); err != nil {
				fmt.Printf("❌ Failed to read migration: %v\n", err)
				return
			}
			fmt.Printf("  ✅ %s at %s\n", name, appliedAt)
		}
		if err := rows.Err(); err != nil {
			fmt.Printf("❌ Rows error: %v\n", err)
			return
		}

	case "--help":
		printHelp()

	default:
		fmt.Println("❓ Unknown command:", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Printf(`📘 Available commands:
  --start             Start the application and run all migrations
	--db-check          Check connection database
  --db-create         Create all database tables from the migration files
  --db-drop [folder]  Drop tables and remove migration records (e.g., --db-drop 03_middleware)
  --status            Show migrations status
  --help              Show this help message
`)
}
