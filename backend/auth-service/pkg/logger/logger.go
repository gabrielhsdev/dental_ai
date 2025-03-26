package logger

// Logger defines the interface for the logging pkg system
type Logger interface {
	Info(msg string, fields ...LogField)
	Warning(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	Debug(msg string, fields ...LogField)
	Sync() error

	// LogObject is used for loggin entire json structures
	LogObject(msg string, obj interface{})
}

type LogField struct {
	Key   string
	Value interface{}
}
