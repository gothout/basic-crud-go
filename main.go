// @title           Basic Crud
// @version         1.0
// @description     API administrativa generica para um Basic Crud

package main

import (
	"basic-crud-go/cmd/cli"
	cmdEnv "basic-crud-go/cmd/configuration/env"
)

func main() {
	// Validate envs
	cmdEnv.ValidateEnvs()
	// Handle CLI commands (e.g., --start, --db-create, --db-drop)
	cli.HandleCLI()
}
