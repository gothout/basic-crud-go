package log

import (
	"fmt"
	"os"
	"strconv"
)

// GetLogLevel retrieves the LOG_LEVEL environment variable.
func GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

// GetLogDatabase retrieves the LOG_DATABASE environment variable.
func GetLogDatabase() string {
	return os.Getenv("LOG_DATABASE")
}

// ValidateLogsEnv ensures all required environment variables for logs are valid.
func ValidateLogsEnv() error {
	logLevel := GetLogLevel()
	if logLevel == "" {
		return fmt.Errorf("environment variable LOG_LEVEL not set")
	}

	levelInt, err := strconv.Atoi(logLevel)
	if err != nil || levelInt > 2 {
		return fmt.Errorf("LOG_LEVEL must be a valid integer <= 2")
	}

	logDatabase := GetLogDatabase()
	if logDatabase == "" {
		return fmt.Errorf("environment variable LOG_DATABASE not set")
	}
	if logDatabase != "TRUE" && logDatabase != "FALSE" {
		return fmt.Errorf("LOG_DATABASE must be either 'TRUE' or 'FALSE'")
	}

	return nil
}
