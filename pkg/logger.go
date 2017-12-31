package logger

import (
	lf "github.com/sniperkit/xlogger/pkg/fields"
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	DebugWithFields(fields lf.Fields, format string, args ...interface{})
	InfoWithFields(fields lf.Fields, format string, args ...interface{})
	WarningWithFields(fields lf.Fields, format string, args ...interface{})
	ErrorWithFields(fields lf.Fields, format string, args ...interface{})
	FatalWithFields(fields lf.Fields, format string, args ...interface{})
	Sync() error
}
