package logger

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	DebugWithFields(fields Fields, format string, args ...interface{})
	InfoWithFields(fields Fields, format string, args ...interface{})
	WarningWithFields(fields Fields, format string, args ...interface{})
	ErrorWithFields(fields Fields, format string, args ...interface{})
	FatalWithFields(fields Fields, format string, args ...interface{})
	Sync() error
}
