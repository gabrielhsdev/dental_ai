package logger

func NewLogField(key string, value interface{}) LogField {
	return LogField{Key: key, Value: value}
}

func LogError(err error) LogField {
	return NewLogField("error", err.Error())
}

func LogInfo(msg string) LogField {
	return NewLogField("message", msg)
}

func LogWarning(msg string) LogField {
	return NewLogField("warning", msg)
}

func LogRequestID(requestId string) LogField {
	return NewLogField("request_id", requestId)
}

func LogUserID(userID string) LogField {
	return NewLogField("user_id", userID)
}

func LogTraceID(traceID string) LogField {
	return NewLogField("trace_id", traceID)
}
