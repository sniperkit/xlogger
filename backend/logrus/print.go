package logrus

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

func (l *Logger) Panic(format string, args ...interface{}) {
	l.decorateEntry(nil).Panic(args...)
}

func (l *Logger) Print(format string, args ...interface{}) {
	l.decorateEntry(nil).Print(args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.decorateEntry(nil).Debug(args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.decorateEntry(nil).Info(args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.decorateEntry(nil).Warning(args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.decorateEntry(nil).Error(args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.decorateEntry(nil).Fatal(args...)
}

func (l *Logger) PrintWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Print(args...)
}

func (l *Logger) PanicWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Panic(args...)
}

func (l *Logger) DebugWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Debug(args...)
}

func (l *Logger) InfoWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Info(args...)
}

func (l *Logger) WarnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Warning(args...)
}

func (l *Logger) ErrorWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Error(args...)
}

func (l *Logger) FatalWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Fatal(args...)
}
