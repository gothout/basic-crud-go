package environment

import (
	"fmt"
	"os"
)

// GetENVIRONMENT retrieves the ENV environment variable.
func GetEnvironment() string {
	return os.Getenv("ENV")
}

// ValidateEnvironmentEnv ensures all required environment variables are set.
func ValidateEnvironmentEnv() error {
	if GetEnvironment() == "" {
		return fmt.Errorf("environment variable ENV not set")
	}
	return nil
}
