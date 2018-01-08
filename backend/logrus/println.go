package logrus

import (
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

func (l *Logger) Panicln(format string, args ...interface{}) {
	l.decorateEntry(nil).Panicln(args...)
}

func (l *Logger) Println(format string, args ...interface{}) {
	l.decorateEntry(nil).Println(args...)
}

func (l *Logger) Debugln(format string, args ...interface{}) {
	l.decorateEntry(nil).Debugln(args...)
}

func (l *Logger) Infoln(format string, args ...interface{}) {
	l.decorateEntry(nil).Infoln(args...)
}

func (l *Logger) Warnln(format string, args ...interface{}) {
	l.decorateEntry(nil).Warningln(args...)
}

func (l *Logger) Errorln(format string, args ...interface{}) {
	l.decorateEntry(nil).Errorln(args...)
}

func (l *Logger) Fatalln(format string, args ...interface{}) {
	l.decorateEntry(nil).Fatalln(args...)
}

func (l *Logger) PrintlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Println(args...)
}

func (l *Logger) PaniclnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Panicln(args...)
}

func (l *Logger) DebuglnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Debugln(args...)
}

func (l *Logger) InfolnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Infoln(args...)
}

func (l *Logger) WarnlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Warningln(args...)
}

func (l *Logger) ErrorlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Errorln(args...)
}

func (l *Logger) FatallnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Fatalln(args...)
}
