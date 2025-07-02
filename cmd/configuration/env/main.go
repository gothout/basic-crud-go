package env

import (
	env "basic-crud-go/internal/configuration/env"
	"log"
)

func ValidateEnvs() {
	if err := env.CheckEnvs(); err != nil {
		log.Fatal(err)
	}
}
