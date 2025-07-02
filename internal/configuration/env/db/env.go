package db

import (
	"fmt"
	"os"
)

// GetDatabaseURL retrieves the DATABASE_URL environment variable.
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

// GetDatabasePort retrieves the DATABASE_PORT environment variable.
func GetDatabasePort() string {
	return os.Getenv("DATABASE_PORT")
}

// GetDatabaseUser retrieves the DATABASE_USER environment variable.
func GetDatabaseUser() string {
	return os.Getenv("DATABASE_USER")
}

// GetDatabasePassword retrieves the DATABASE_PW environment variable.
func GetDatabasePassword() string {
	return os.Getenv("DATABASE_PW")
}

// GetDatabaseName retrieves the DATABASE_NAME environment variable.
func GetDatabaseName() string {
	return os.Getenv("DATABASE_NAME")
}

// GetDatabaseSSL retrieves the DATABASE_SSL environment variable.
func GetDatabaseSSL() string {
	return os.Getenv("DATABASE_SSL")
}

// ValidateDatabaseEnv ensures all required database environment variables are set.
func ValidateDatabaseEnv() error {
	if GetDatabaseURL() == "" {
		return fmt.Errorf("environment variable DATABASE_URL not set")
	}
	if GetDatabasePort() == "" {
		return fmt.Errorf("environment variable DATABASE_PORT not set")
	}
	if GetDatabaseUser() == "" {
		return fmt.Errorf("environment variable DATABASE_USER not set")
	}
	if GetDatabasePassword() == "" {
		return fmt.Errorf("environment variable DATABASE_PW not set")
	}
	if GetDatabaseName() == "" {
		return fmt.Errorf("environment variable DATABASE_NAME not set")
	}
	if GetDatabaseSSL() != "disable" && GetDatabaseSSL() != "enabled" {
		return fmt.Errorf("environment variable DATABASE_SSL must be 'disable' or 'enabled'")
	}
	return nil
}
