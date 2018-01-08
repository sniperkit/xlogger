package zap

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

func (l *Logger) Println(format string, args ...interface{}) { l.logger.Infow(format, args...) }
func (l *Logger) Panicln(format string, args ...interface{}) { l.logger.Panicw(format, args...) }
func (l *Logger) Fatalln(format string, args ...interface{}) { l.logger.Fatalw(format, args...) }
func (l *Logger) Debugln(format string, args ...interface{}) { l.logger.Debugw(format, args...) }
func (l *Logger) Infoln(format string, args ...interface{})  { l.logger.Infow(format, args...) }
func (l *Logger) Warnln(format string, args ...interface{})  { l.logger.Warnw(format, args...) }
func (l *Logger) Errorln(format string, args ...interface{}) { l.logger.Errorw(format, args...) }

func (l *Logger) PrintlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Infow(format, args...)
}

func (l *Logger) PaniclnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Panicw(format, args...)
}

func (l *Logger) DebuglnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Debugw(format, args...)
}

func (l *Logger) InfolnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Infow(format, args...)
}

func (l *Logger) WarnlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Warnw(format, args...)
}

func (l *Logger) ErrorlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Errorw(format, args...)
}

func (l *Logger) FatallnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Fatalw(format, args...)
}
