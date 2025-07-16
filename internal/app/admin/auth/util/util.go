package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateSecureToken(byteLength int) (string, error) {
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate secure token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
