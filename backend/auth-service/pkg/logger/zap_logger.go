package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	Zap *zap.Logger
}

func (z *ZapLogger) Info(msg string, fields ...LogField) {
	z.Zap.Info(msg, convertFields(fields)...)
}

func (z *ZapLogger) Warning(msg string, fields ...LogField) {
	z.Zap.Warn(msg, convertFields(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...LogField) {
	z.Zap.Error(msg, convertFields(fields)...)
}

func (z *ZapLogger) Debug(msg string, fields ...LogField) {
	z.Zap.Debug(msg, convertFields(fields)...)
}

func (z *ZapLogger) LogObject(msg string, obj interface{}) {
	z.Zap.Info(msg, zap.Any("data", obj))
}

func (z *ZapLogger) Sync() error {
	return z.Zap.Sync()
}

func convertFields(fields []LogField) []zap.Field {
	var zapFields []zap.Field
	for _, field := range fields {
		zapFields = append(zapFields, zap.Any(field.Key, field.Value))
	}
	return zapFields
}
