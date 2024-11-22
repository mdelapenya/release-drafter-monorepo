package main

import (
	"log"
	"os"
)

// Logger interface defines the methods for logging
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

// DefaultLogger is the default implementation of the Logger interface
type DefaultLogger struct {
	logger *log.Logger
}

// NewDefaultLogger creates a new DefaultLogger
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *DefaultLogger) Debug(msg string) {
	l.log("DEBUG", msg)
}

func (l *DefaultLogger) Info(msg string) {
	l.log("INFO", msg)
}

func (l *DefaultLogger) Warn(msg string) {
	l.log("WARN", msg)
}

func (l *DefaultLogger) Error(msg string) {
	l.log("ERROR", msg)
}

func (l *DefaultLogger) Fatal(msg string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Fatalln(msg)
}

func (l *DefaultLogger) log(prefix string, msg string) {
	l.logger.SetPrefix(prefix + ": ")
	l.logger.Println(msg)
}
