package logger

import (
	"log"
)

type stackedCustomLogger struct {
	prefix string
}

func NewStackedCustomLogger(prefix string) Logger {
	return &stackedCustomLogger{prefix: prefix}
}

// Info prints the message in yellow to the console with a timestamp
func (l *stackedCustomLogger) Info(msg string) {
	log.Printf("[%s] [INFO] \033[33m%s\033[0m\n", l.prefix, msg)
}

// // Error prints the error message in red to the console with a timestamp
func (l *stackedCustomLogger) Error(err error) {
	log.Printf("[%s] [ERROR] \033[31mERROR: %s\033[0m\n", l.prefix, err.Error())
}

// Message prints the message in green to the console with a timestamp
func (l *stackedCustomLogger) Message(msg string) {
	log.Printf("[%s] [MESSAGE] \033[32m%s\033[0m\n", l.prefix, msg)
}

// Debug prints the message in blue to the console with a timestamp
func (l *stackedCustomLogger) Debug(msg string) {
	log.Printf("[%s] [DEBUG] \033[34m%s\033[0m\n", l.prefix, msg)
}
