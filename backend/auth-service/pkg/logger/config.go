package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoadLogger() (Logger, error) {
	// Set up the Zap production configuration with context-specific formats.
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Optionally, add fields for context like environment, service name, etc.
	// You can enrich logs by adding custom fields here if required.
	// config.AddCallerSkip(1) // To skip one level for stack trace information (customize as needed).

	// If we are writing logs to a file, you can set the log file path here.
	// For example, add this if using `file` output:
	// config.OutputPaths = []string{"stdout", "/var/log/service.log"}

	// Build the logger instance.
	zapLogger, err := config.Build(zap.AddCaller())
	if err != nil {
		return nil, err
	}

	// Build the logger instance.
	return &ZapLogger{Zap: zapLogger}, nil
}
