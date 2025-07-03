package postgres

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() {
	db := GetDB()

	// Cria tabela de controle de migrations
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations_applied (
			id SERIAL PRIMARY KEY,
			filename TEXT UNIQUE,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatalf("❌ Erro ao criar tabela migrations_applied: %v", err)
	}

	// Lê todos os arquivos .sql da pasta migrate/
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
		log.Fatalf("❌ Erro ao ler diretório de migrations: %v", err)
	}

	// Ordena os arquivos pelo nome (respeita prefixos 01_, 02_, etc.)
	sort.Strings(migrations)

	for _, path := range migrations {
		filename := filepath.Base(path)

		var alreadyApplied bool
		err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM migrations_applied WHERE filename = $1)`, filename).Scan(&alreadyApplied)
		if err != nil {
			log.Fatalf("❌ Erro ao verificar migration %s: %v", filename, err)
		}

		if alreadyApplied {
			log.Printf("🔸 Migration já aplicada: %s", filename)
			continue
		}

		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("❌ Erro ao ler arquivo %s: %v", path, err)
		}

		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			log.Fatalf("❌ Erro ao executar migration %s: %v", filename, err)
		}

		_, err = db.Exec(`INSERT INTO migrations_applied (filename) VALUES ($1)`, filename)
		if err != nil {
			log.Fatalf("❌ Erro ao registrar migration %s: %v", filename, err)
		}

		log.Printf("✅ Migration aplicada: %s", filename)
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
		log.Fatalf("❌ Nenhuma tabela mapeada para a pasta %s", folder)
	}

	// Apaga as tabelas (ordem reversa para respeitar dependências)
	for i := len(tables) - 1; i >= 0; i-- {
		table := tables[i]
		query := "DROP TABLE IF EXISTS " + table + " CASCADE"
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("❌ Erro ao deletar tabela %s: %v", table, err)
		} else {
			log.Printf("🗑️  Tabela %s deletada com sucesso", table)
		}
	}

	// Remove registros de migrations_applied referentes a arquivos da pasta
	files, err := filepath.Glob("cmd/migrate/" + folder + "/*.sql")
	if err != nil {
		log.Fatalf("❌ Erro ao buscar arquivos da pasta %s: %v", folder, err)
	}

	for _, path := range files {
		filename := filepath.Base(path)
		_, err := db.Exec(`DELETE FROM migrations_applied WHERE filename = $1`, filename)
		if err != nil {
			log.Printf("⚠️  Erro ao remover migration %s do histórico: %v", filename, err)
		} else {
			log.Printf("🧹 Migration %s removida do histórico", filename)
		}
	}
}
