package main

import (
	cmdEnv "basic-crud-go/cmd/configuration/env"
)

func main() {
	// Validate envs
	cmdEnv.ValidateEnvs()
}
