package postgres

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() {
	db := GetDB()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations_applied (
			id SERIAL PRIMARY KEY,
			filename TEXT UNIQUE,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatalf("❌ Error creating migrations_applied table: %v", err)
	}

	migrations := []string{}
	root := "cmd/migrate"
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".sql") {
			migrations = append(migrations, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("❌ Error reading migrations directory: %v", err)
	}

	// Ordena por data contida no nome do arquivo (ascendente)
	sort.Slice(migrations, func(i, j int) bool {
		return extractDate(migrations[i]) < extractDate(migrations[j])
	})

	for _, path := range migrations {
		filename := filepath.Base(path)

		var alreadyApplied bool
		err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM migrations_applied WHERE filename = $1)`, filename).Scan(&alreadyApplied)
		if err != nil {
			log.Fatalf("❌ Error checking migration %s: %v", filename, err)
		}

		if alreadyApplied {
			log.Printf("🔸 Migration already applied: %s", filename)
			continue
		}

		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("❌ Error reading file %s: %v", path, err)
		}

		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			log.Fatalf("❌ Error executing migration %s: %v", filename, err)
		}

		_, err = db.Exec(`INSERT INTO migrations_applied (filename) VALUES ($1)`, filename)
		if err != nil {
			log.Fatalf("❌ Error inserting migration record %s: %v", filename, err)
		}

		log.Printf("✅ Migration applied: %s", filename)
	}
}

var rollbackMapByFolder = map[string][]string{
	"01_enterprise": {"enterprise"},
	"02_user":       {"user"},
	"03_middleware": {"user_permission", "permission", "action", "module"},
}

func RollbackByFolder(folder string) {
	db := GetDB()

	tables, ok := rollbackMapByFolder[folder]
	if !ok {
		log.Fatalf("❌ No table mapping found for folder: %s", folder)
	}

	// Deleta os arquivos em ordem de data DESC (mais recentes primeiro)
	files, err := filepath.Glob("cmd/migrate/" + folder + "/*.sql")
	if err != nil {
		log.Fatalf("❌ Error scanning folder %s: %v", folder, err)
	}

	sort.Slice(files, func(i, j int) bool {
		return extractDate(files[i]) > extractDate(files[j])
	})

	// Drop tables (in reverse order of dependency)
	for i := len(tables) - 1; i >= 0; i-- {
		table := tables[i]
		query := fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table)
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("❌ Error dropping table %s: %v", table, err)
		} else {
			log.Printf("🗑️  Table %s dropped successfully", table)
		}
	}

	// Clean up migration records for that folder
	for _, path := range files {
		filename := filepath.Base(path)
		_, err := db.Exec(`DELETE FROM migrations_applied WHERE filename = $1`, filename)
		if err != nil {
			log.Printf("⚠️  Failed to delete migration record %s: %v", filename, err)
		} else {
			log.Printf("🧹 Migration %s removed from history", filename)
		}
	}
}

// extractDate extracts a comparable date string from the filename (e.g. _20250703 -> 20250703)
func extractDate(path string) string {
	filename := filepath.Base(path)
	parts := strings.Split(filename, "_")
	if len(parts) < 2 {
		return ""
	}
	dateWithExt := parts[len(parts)-1]
	return strings.TrimSuffix(dateWithExt, filepath.Ext(dateWithExt))
}
