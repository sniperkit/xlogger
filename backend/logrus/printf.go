package logrus

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.decorateEntry(nil).Panicf(format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.decorateEntry(nil).Printf(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.decorateEntry(nil).Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.decorateEntry(nil).Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.decorateEntry(nil).Warningf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.decorateEntry(nil).Errorf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.decorateEntry(nil).Fatalf(format, args...)
}

func (l *Logger) PrintfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Printf(format, args...)
}

func (l *Logger) PanicfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Panicf(format, args...)
}

func (l *Logger) DebugfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Debugf(format, args...)
}

func (l *Logger) InfofWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Infof(format, args...)
}

func (l *Logger) WarnfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Warningf(format, args...)
}

func (l *Logger) ErrorfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Errorf(format, args...)
}

func (l *Logger) FatalfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Fatalf(format, args...)
}
