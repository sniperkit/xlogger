package log

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

// Arguments are handled in the manner of fmt.Printf
func (l *Logger) Panicf(format string, args ...interface{}) {}
func (l *Logger) Printf(format string, args ...interface{}) {}
func (l *Logger) Debugf(format string, args ...interface{}) {}
func (l *Logger) Infof(format string, args ...interface{})  {}
func (l *Logger) Warnf(format string, args ...interface{})  {}
func (l *Logger) Errorf(format string, args ...interface{}) {}
func (l *Logger) Fatalf(format string, args ...interface{}) {}

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (l *Logger) PrintfWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) PanicfWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) DebugfWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) InfofWithFields(fields lfi.Fields, format string, args ...interface{})  {}
func (l *Logger) WarnfWithFields(fields lfi.Fields, format string, args ...interface{})  {}
func (l *Logger) ErrorfWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) FatalfWithFields(fields lfi.Fields, format string, args ...interface{}) {}
