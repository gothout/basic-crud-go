// @title           Basic CRUD API
// @version         1.0
// @description     Generic administrative API for a Basic Crud
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter the token in the format: Bearer <your_token>

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
