package postgres

import (
	env "basic-crud-go/internal/configuration/env/db"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	db   *sql.DB
	once sync.Once
)

// InitPostgres initializes the database connection
func InitPostgres() *sql.DB {
	once.Do(func() {
		dsn := buildDSN()
		var err error

		db, err = sql.Open("pgx", dsn)
		if err != nil {
			log.Fatalf("error open connection to the database: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("error test connection database: %v", err)
		}

		log.Println("sucess connect database postgresql")
	})

	return db
}

// GetDB returns the existing DB connection instance.
// It panics if InitPostgres was not called before.
func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("database connection is not initialized. Call InitPostgres() first.")
	}
	return db
}

func buildDSN() string {
	host := env.GetDatabaseURL()
	port := env.GetDatabasePort()
	user := env.GetDatabaseUser()
	pass := env.GetDatabasePassword()
	name := env.GetDatabaseName()
	ssl := env.GetDatabaseSSL()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)
}
