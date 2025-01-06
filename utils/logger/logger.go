package logger

// Logger defines the interface for our custom logger
type Logger interface {
	Info(msg string)
	Message(msg string)
	Debug(msg string)
	Error(err error)
}
