package log

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(format string, args ...interface{}) {}
func (l *Logger) Panic(format string, args ...interface{}) {}
func (l *Logger) Debug(format string, args ...interface{}) {}
func (l *Logger) Info(format string, args ...interface{})  {}
func (l *Logger) Warn(format string, args ...interface{})  {}
func (l *Logger) Error(format string, args ...interface{}) {}
func (l *Logger) Fatal(format string, args ...interface{}) {}

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (l *Logger) PrintWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) PanicWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) DebugWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) InfoWithFields(fields lfi.Fields, format string, args ...interface{})  {}
func (l *Logger) WarnWithFields(fields lfi.Fields, format string, args ...interface{})  {}
func (l *Logger) ErrorWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) FatalWithFields(fields lfi.Fields, format string, args ...interface{}) {}
