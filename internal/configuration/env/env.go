package env

import (
	"basic-crud-go/internal/configuration/env/db"
	envEnvs "basic-crud-go/internal/configuration/env/enviroument"
	logEnvs "basic-crud-go/internal/configuration/env/log"
	securityEnvs "basic-crud-go/internal/configuration/env/security"
	serverEnvs "basic-crud-go/internal/configuration/env/server"
	"fmt"

	"github.com/joho/godotenv"
)

// CheckEnvs checks if the .env file exists and returns an appropriate error if it doesn't.
func CheckEnvs() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// Validate environment variable principal
	err = envEnvs.ValidateEnvironmentEnv()
	if err != nil {
		return err
	}
	// Validate server envs
	err = serverEnvs.ValidateServerEnv()
	if err != nil {
		return err
	}
	// Validate log server envs
	err = logEnvs.ValidateLogsEnv()
	if err != nil {
		return err
	}
	// Validate security envs
	err = securityEnvs.ValidateSecurityEnv()
	if err != nil {
		return err
	}
	// Validate database envsserverEnvs "basic-crud-go/internal/configuration/env/server"
	err = db.ValidateDatabaseEnv()
	if err != nil {
		return err
	}

	return nil
}
