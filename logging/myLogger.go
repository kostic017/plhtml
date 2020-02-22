package logging

import (
	"fmt"
	"log"
	"os"
)

type logLevel int

const (
	// Off to turn off logging
	Off logLevel = iota
	// Fine for fine-grained information
	Fine logLevel = iota
	// Debug for diagnostic information
	Debug logLevel = iota
	// Info for normal application behavior
	Info logLevel = iota
	// Warn for potentially harmful situations
	Warn logLevel = iota
	// Error for not so severe errors
	Error logLevel = iota
	// Fatal for very severe errors
	Fatal logLevel = iota
	// All levels
	All logLevel = iota
)

// MyLogger wraps log.Logger and adds support for logging levels
type MyLogger struct {
	level     logLevel
	stdLogger *log.Logger
}

// New creates a new logger
func New(name string) *MyLogger {
	logger := new(MyLogger)
	logger.level = Info
	logger.stdLogger = log.New(os.Stdout, name, 0)
	return logger
}

// SetLevel to change logging level
func (logger *MyLogger) SetLevel(level logLevel) {
	logger.level = level
}

func (logger MyLogger) log(prefix string, format string, args ...interface{}) {
	logger.stdLogger.Print(fmt.Sprintf(" %-5s %s", prefix, fmt.Sprintf(format, args...)))
}

// Fine for fine-grained information
func (logger MyLogger) Fine(format string, args ...interface{}) {
	if Fine >= logger.level {
		logger.log("FINE", format, args...)
	}
}

// Debug for diagnostic information
func (logger MyLogger) Debug(format string, args ...interface{}) {
	if Debug >= logger.level {
		logger.log("DEBUG", format, args...)
	}
}

// Info for normal application behavior
func (logger MyLogger) Info(format string, args ...interface{}) {
	if Info >= logger.level {
		logger.log("INFO", format, args...)
	}
}

// Warn for potentially harmful situations
func (logger MyLogger) Warn(format string, args ...interface{}) {
	if Warn >= logger.level {
		logger.log("WARN", format, args...)
	}
}

// Error for not so severe errors
func (logger MyLogger) Error(format string, args ...interface{}) {
	if Error >= logger.level {
		logger.log("ERROR", format, args...)
	}
}

// Fatal for very severe errors
func (logger MyLogger) Fatal(format string, args ...interface{}) {
	if Fatal >= logger.level {
		logger.log("FATAL", format, args...)
	}
}
