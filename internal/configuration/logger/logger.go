// Package logger provides a simple and extensible logging utility for writing
// application logs to separate daily files based on severity level.
//
// Log messages are stored in the "logs/" directory, with filenames in the format:
//
//	info_02072025.log, warning_02072025.log, error_02072025.log
//
// Each log entry includes:
//   - log level
//   - module name
//   - function name
//   - timestamp
//   - message content
//
// The logger supports generic messages of any type, including strings, errors, and structs.
// Structs will be automatically marshaled into JSON format for readability.
//
// Log level filtering is controlled via the environment variable LOG_LEVEL:
//
//	0 = log all
//	1 = log only warning, info, and error
//	2 = log only info and error
//	3 = log only error
//
// There are two logging functions:
//
//	Log(level, module, function, message)
//	LogWithAutoFuncName(level, module, message)
//
// Use Log for explicit function naming, or LogWithAutoFuncName for automatic detection
// of the calling function (note: slightly slower due to runtime inspection).
//
// Examples:
//
//	logger.Log(logger.Info, "UserService", "CreateUser", "User created successfully")
//	logger.LogWithAutoFuncName(logger.Warning, "AuthMiddleware", fmt.Errorf("token expired"))
//	logger.LogWithAutoFuncName(logger.Error, "DatabaseService", struct {
//	    Code int    `json:"code"`
//	    Msg  string `json:"msg"`
//	}{Code: 500, Msg: "Database connection failed"})
package logger

import (
	env "basic-crud-go/internal/configuration/env/log"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

var currentLogLevel = -1

// Log writes a message to the corresponding log file based on the level.
// It supports any data type for the message (string, struct, error, etc.)
func Log[T any](level LogLevel, module, function string, message T) {
	if currentLogLevel == -1 {
		currentLogLevel = loadLogLevel()
	}
	// If LOG_LEVEL=0 (all), allow everything
	// Otherwise, skip if current level is more restrictive than message
	if currentLogLevel > 0 && levelPriority(level) < currentLogLevel {
		return
	}

	now := time.Now()
	date := now.Format("02012006") // e.g., 01072025
	timestamp := now.Format("2006-01-02 15:04:05")

	// Convert message to string (JSON for structs, etc.)
	msgString := convertToString(message)

	logMsg := fmt.Sprintf("[%s] [%s] [%s] %s - %s\n", levelString(level), module, function, timestamp, msgString)

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

// LogWithAutoFuncName logs the message using the function name of the caller automatically.
// This is slightly slower than manually providing the function name.
func LogWithAutoFuncName[T any](level LogLevel, module string, message T) {
	function := getCallerFunctionName()
	Log(level, module, function, message)
}

// getCallerFunctionName extracts the name of the function that called LogWithAutoFuncName.
func getCallerFunctionName() string {
	pc, _, _, ok := runtime.Caller(2) // 2 levels up: caller of LogWithAutoFuncName
	if !ok {
		return "unknown"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	fullName := fn.Name() // e.g., "basic-crud-go/internal/app/.../service.(*enterpriseServiceImpl).ReadAllEnterprise"
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
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

// levelPriority returns a numeric priority for comparison
// Lower number = less severe
func levelPriority(level LogLevel) int {
	switch level {
	case Error:
		return 3
	case Info:
		return 2
	case Warning:
		return 1
	default:
		return 0
	}
}

// loadLogLevel reads the LOG_LEVEL environment variable and returns the active threshold
func loadLogLevel() int {
	val := env.GetLogLevel()
	switch val {
	case "0":
		return 0 // log everything
	case "1":
		return 1 // warning, info, error
	case "2":
		return 2 // info, error
	case "3":
		return 3 // error only
	default:
		return 0 // fallback to all
	}
}
