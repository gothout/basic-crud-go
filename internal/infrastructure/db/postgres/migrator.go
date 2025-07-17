package postgres

import (
	"fmt"
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

	root := "cmd/migrate"
	subdirs, err := os.ReadDir(root)
	if err != nil {
		log.Fatalf("❌ Error reading root migrate directory: %v", err)
	}

	// Ordena as pastas numericamente: 01_, 02_, etc.
	sort.Slice(subdirs, func(i, j int) bool {
		return subdirs[i].Name() < subdirs[j].Name()
	})

	for _, dir := range subdirs {
		if !dir.IsDir() {
			continue
		}

		dirPath := filepath.Join(root, dir.Name())
		files, err := os.ReadDir(dirPath)
		if err != nil {
			log.Printf("⚠️  Error reading dir %s: %v", dirPath, err)
			continue
		}

		sqlFiles := []string{}
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
				sqlFiles = append(sqlFiles, filepath.Join(dirPath, file.Name()))
			}
		}

		// Ordena arquivos dentro da pasta por data extraída do nome
		sort.Slice(sqlFiles, func(i, j int) bool {
			return extractDate(sqlFiles[i]) < extractDate(sqlFiles[j])
		})

		for _, path := range sqlFiles {
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
}

// 🔁 Rollback por pasta com base no nome da pasta (ex: 02_user)
var rollbackMapByFolder = map[string][]string{
	"01_enterprise": {"enterprise"},
	"02_user":       {"user", "admin_token"},
	"03_permission": {"admin_permission", "user_permission"},
}

func RollbackByFolder(folder string) {
	db := GetDB()

	tables, ok := rollbackMapByFolder[folder]
	if !ok {
		log.Fatalf("❌ No table mapping found for folder: %s", folder)
	}

	files, err := filepath.Glob("cmd/migrate/" + folder + "/*.sql")
	if err != nil {
		log.Fatalf("❌ Error scanning folder %s: %v", folder, err)
	}

	sort.Slice(files, func(i, j int) bool {
		return extractDate(files[i]) > extractDate(files[j])
	})

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

// 🔻 Drop completo de todas as tabelas
func DropAllMigrations() {
	db := GetDB()

	log.Println("🧨 Dropping all migration tables in reverse order...")

	dropOrder := []string{
		"user_permission",
		"admin_permission",
		"user",
		"enterprise",
	}

	for _, table := range dropOrder {
		query := fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table)
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("❌ Error dropping table %s: %v", table, err)
		} else {
			log.Printf("🗑️  Table %s dropped successfully", table)
		}
	}

	_, err := db.Exec(`DROP TABLE IF EXISTS migrations_applied CASCADE`)
	if err != nil {
		log.Printf("⚠️  Error dropping migrations_applied: %v", err)
	} else {
		log.Println("🧹 Table migrations_applied dropped successfully")
	}
}

// 🔍 Extrai a data do nome do arquivo
func extractDate(path string) string {
	filename := filepath.Base(path)
	parts := strings.Split(filename, "_")
	if len(parts) < 2 {
		return ""
	}
	dateWithExt := parts[len(parts)-1]
	return strings.TrimSuffix(dateWithExt, filepath.Ext(dateWithExt))
}
