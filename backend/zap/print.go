package zap

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

func (l *Logger) Print(format string, args ...interface{}) { l.logger.Info(args...) }
func (l *Logger) Debug(format string, args ...interface{}) { l.logger.Debug(args...) }
func (l *Logger) Info(format string, args ...interface{})  { l.logger.Info(args...) }
func (l *Logger) Warn(format string, args ...interface{})  { l.logger.Warn(args...) }
func (l *Logger) Error(format string, args ...interface{}) { l.logger.Error(args...) }
func (l *Logger) Fatal(format string, args ...interface{}) { l.logger.Fatal(args...) }
func (l *Logger) Panic(format string, args ...interface{}) { l.logger.Panic(args...) }

func (l *Logger) PrintWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Info(args...)
}

func (l *Logger) PanicWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Panic(args...)
}

func (l *Logger) DebugWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Debug(args...)
}

func (l *Logger) InfoWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Info(args...)
}

func (l *Logger) WarnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Warn(args...)
}

func (l *Logger) ErrorWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Error(args...)
}

func (l *Logger) FatalWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Fatal(args...)
}
