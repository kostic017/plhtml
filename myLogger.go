package main

import (
	"fmt"
	"log"
	"os"
)

type logLevel int

const (
	lvlOff   logLevel = iota
	lvlFine  logLevel = iota
	lvlDebug logLevel = iota
	lvlInfo  logLevel = iota
	lvlWarn  logLevel = iota
	lvlError logLevel = iota
	lvlFatal logLevel = iota
	lvlAll   logLevel = iota
)

type myLogger struct {
	level     logLevel
	stdLogger *log.Logger
}

func newLogger(name string) *myLogger {
	logger := new(myLogger)
	logger.level = lvlInfo
	logger.stdLogger = log.New(os.Stdout, name, 0)
	return logger
}

func (logger *myLogger) setLevel(level logLevel) {
	logger.level = level
}

func (logger myLogger) log(prefix string, format string, args ...interface{}) {
	logger.stdLogger.Print(fmt.Sprintf(" %-5s %s", prefix, fmt.Sprintf(format, args...)))
}

func (logger myLogger) fine(format string, args ...interface{}) {
	if lvlFine >= logger.level {
		logger.log("FINE", format, args...)
	}
}

func (logger myLogger) debug(format string, args ...interface{}) {
	if lvlDebug >= logger.level {
		logger.log("DEBUG", format, args...)
	}
}

func (logger myLogger) info(format string, args ...interface{}) {
	if lvlInfo >= logger.level {
		logger.log("INFO", format, args...)
	}
}

func (logger myLogger) warn(format string, args ...interface{}) {
	if lvlWarn >= logger.level {
		logger.log("WARN", format, args...)
	}
}

func (logger myLogger) error(format string, args ...interface{}) {
	if lvlError >= logger.level {
		logger.log("ERROR", format, args...)
	}
}

func (logger myLogger) fatal(format string, args ...interface{}) {
	if lvlFatal >= logger.level {
		logger.log("FATAL", format, args...)
	}
}
