// @title           Basic Crud
// @version         1.0
// @description     API administrativa generica para um Basic Crud

package main

import (
	cmdEnv "basic-crud-go/cmd/configuration/env"
	cmdServer "basic-crud-go/cmd/server"
	"basic-crud-go/internal/infrastructure/db/postgres"
)

func main() {
	// Validate envs
	cmdEnv.ValidateEnvs()

	// Inicialize connection database
	postgres.InitPostgres()

	// Inicialize server
	router := cmdServer.InitServer()
	cmdServer.StartServer(router)
}
