package zap

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

func (l *Logger) Printf(format string, args ...interface{}) { l.logger.Infof(format, args...) }
func (l *Logger) Debugf(format string, args ...interface{}) { l.logger.Debugf(format, args...) }
func (l *Logger) Infof(format string, args ...interface{})  { l.logger.Infof(format, args...) }
func (l *Logger) Warnf(format string, args ...interface{})  { l.logger.Warnf(format, args...) }
func (l *Logger) Errorf(format string, args ...interface{}) { l.logger.Errorf(format, args...) }
func (l *Logger) Fatalf(format string, args ...interface{}) { l.logger.Fatalf(format, args...) }
func (l *Logger) Panicf(format string, args ...interface{}) { l.logger.Panicf(format, args...) }

func (l *Logger) PrintfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Infof(format, args...)
}

func (l *Logger) PanicfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Panicf(format, args...)
}

func (l *Logger) DebugfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Debugf(format, args...)
}

func (l *Logger) InfofWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Infof(format, args...)
}

func (l *Logger) WarnfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Warnf(format, args...)
}

func (l *Logger) ErrorfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Errorf(format, args...)
}

func (l *Logger) FatalfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Fatalf(format, args...)
}
