// @title           Basic Crud
// @version         1.0
// @description     API administrativa generica para um Basic Crud

package main

import (
	"basic-crud-go/cmd/cli"
	"basic-crud-go/cmd/configuration/env"
)

func main() {
	// Validate envs
	env.ValidateEnvs()
	// Handle CLI commands (e.g., --start, --db-create, --db-drop)
	cli.HandleCLI()
}
