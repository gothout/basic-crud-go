package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const logDir = "logs"

// LogLevel defines supported log levels
type LogLevel string

const (
	Info    LogLevel = "info"
	Warning LogLevel = "warning"
	Error   LogLevel = "error"
)

// Log writes a message to the corresponding log file based on the level.
// It supports any data type for the message (string, struct, error, etc.)
func Log[T any](level LogLevel, module string, message T) {
	now := time.Now()
	date := now.Format("02012006") // e.g., 01072025
	timestamp := now.Format("2006-01-02 15:04:05")

	// Convert message to string (JSON for structs, etc.)
	msgString := convertToString(message)

	logMsg := fmt.Sprintf("[%s] [%s] %s - %s\n", levelString(level), module, timestamp, msgString)

	// Check if log directory exists; create only if needed
	if _, statErr := os.Stat(logDir); os.IsNotExist(statErr) {
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			fmt.Printf("Failed to create log directory: %v\n", err)
			return
		}
	}

	// Define log file path by type and date
	filename := fmt.Sprintf("%s_%s.log", string(level), date)
	logPath := filepath.Join(logDir, filename)

	// Open file in append mode (create if not exists)
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return
	}
	defer file.Close()

	// Write the log message
	if _, err := file.WriteString(logMsg); err != nil {
		fmt.Printf("Failed to write to log file: %v\n", err)
	}
}

// convertToString converts any input to a string for logging
func convertToString[T any](input T) string {
	switch v := any(input).(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		// Encode struct as JSON
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("Error encoding JSON: %v", err)
		}
		return string(jsonBytes)
	}
}

// levelString formats the log level string in uppercase
func levelString(level LogLevel) string {
	switch level {
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
