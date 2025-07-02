package main

import (
	cmdEnv "basic-crud-go/cmd/configuration/env"
	"basic-crud-go/internal/infrastructure/db/postgres"
)

func main() {
	// Validate envs
	cmdEnv.ValidateEnvs()

	// Inicialize connection database
	postgres.InitPostgres()
}
