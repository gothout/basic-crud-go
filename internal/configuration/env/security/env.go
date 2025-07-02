package security

import (
	"fmt"
	"os"
)

// GetRecoveryEmail retrieves the RECOVERY_EMAIL environment variable.
func GetRecoveryEmail() string {
	return os.Getenv("RECOVERY_EMAIL")
}

// GetRecoveryPassword retrieves the RECOVERY_PWD environment variable.
func GetRecoveryPassword() string {
	return os.Getenv("RECOVERY_PWD")
}

// ValidateSecurityEnv ensures all required security-related environment variables are set.
func ValidateSecurityEnv() error {
	if GetRecoveryEmail() == "" {
		return fmt.Errorf("environment variable RECOVERY_EMAIL not set")
	}
	if GetRecoveryPassword() == "" {
		return fmt.Errorf("environment variable RECOVERY_PWD not set")
	}
	return nil
}
