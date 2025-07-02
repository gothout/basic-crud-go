package server

import (
	"fmt"
	"os"
)

// GetHTTPPort retrieves the HTTP_PORT environment variable.
func GetHTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

// GetHTTPSPort retrieves the HTTPS_PORT environment variable.
func GetHTTPSPort() string {
	return os.Getenv("HTTPS_PORT")
}

// GetDNS retrieves the DNS environment variable.
func GetDNS() string {
	return os.Getenv("DNS")
}

// ValidateServerEnv ensures all required environment variables are set.
func ValidateServerEnv() error {
	if GetHTTPPort() == "" {
		return fmt.Errorf("environment variable HTTP_PORT not set")
	}
	if GetHTTPSPort() == "" {
		return fmt.Errorf("environment variable HTTPS_PORT not set")
	}
	if GetDNS() == "" {
		return fmt.Errorf("environment variable DNS not set")
	}
	return nil
}
