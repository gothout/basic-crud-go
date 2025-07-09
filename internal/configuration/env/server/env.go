package server

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

// GetListenServer retrieves the LISTEN_SERVER environment variable.
func GetListenServer() string {
	return os.Getenv("LISTEN_SERVER")
}

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

// GetCorsOrigins retrieves the CORS origins for both DEV and PROD environments.
func GetCorsOrigins() (devCors string, prodCors string) {
	return os.Getenv("CORS_DEV"), os.Getenv("CORS_PROD")
}

// GetHTTPSuse checks if HTTPS should be used and if certificates exist and are valid.
func GetHTTPSuse() bool {
	https := strings.ToUpper(os.Getenv("HTTPS"))
	if https != "TRUE" {
		return false
	}

	certPath := "./certificates/cert.crt"
	keyPath := "./certificates/privkey.key"

	// Check if both certificate files exist
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		fmt.Println("Certificate file not found:", certPath)
		return false
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		fmt.Println("Private key file not found:", keyPath)
		return false
	}

	// Try loading the certificate to validate it
	_, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		fmt.Println("Invalid TLS certificate or key:", err)
		return false
	}

	return true
}

// ValidateServerEnv ensures all required environment variables are set.
func ValidateServerEnv() error {
	if GetListenServer() == "" {
		return fmt.Errorf("environment variable LISTEN_SERVER not set")
	}
	if GetHTTPPort() == "" {
		return fmt.Errorf("environment variable HTTP_PORT not set")
	}
	if GetHTTPSPort() == "" {
		return fmt.Errorf("environment variable HTTPS_PORT not set")
	}
	if GetDNS() == "" {
		return fmt.Errorf("environment variable DNS not set")
	}
	devCors, prodCors := GetCorsOrigins()
	if devCors == "" {
		return fmt.Errorf("environment variable CORS_DEV not set")
	}
	if prodCors == "" {
		return fmt.Errorf("environment variable CORS_PROD not set")
	}

	https := strings.ToUpper(os.Getenv("HTTPS"))
	if https != "TRUE" && https != "FALSE" {
		return fmt.Errorf("environment variable HTTPS must be TRUE or FALSE")
	}

	if https == "TRUE" && !GetHTTPSuse() {
		return fmt.Errorf("HTTPS is enabled but certificates are missing or invalid")
	}

	return nil
}
