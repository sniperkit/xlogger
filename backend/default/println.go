package log

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Println(format string, args ...interface{}) {}
func (l *Logger) Panicln(format string, args ...interface{}) {}
func (l *Logger) Debugln(format string, args ...interface{}) {}
func (l *Logger) Infoln(format string, args ...interface{})  {}
func (l *Logger) Warnln(format string, args ...interface{})  {}
func (l *Logger) Errorln(format string, args ...interface{}) {}
func (l *Logger) Fatalln(format string, args ...interface{}) {}

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (l *Logger) PrintlnWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) PaniclnWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) DebuglnWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) InfolnWithFields(fields lfi.Fields, format string, args ...interface{})  {}
func (l *Logger) WarnlnWithFields(fields lfi.Fields, format string, args ...interface{})  {}
func (l *Logger) ErrorlnWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) FatallnWithFields(fields lfi.Fields, format string, args ...interface{}) {}
