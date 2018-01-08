package gomol

import (
	"os"

	"github.com/aphistic/gomol"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

// not working properly yet

func (l *Logger) Print(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Panic(format string, args ...interface{}) {
	l.log(gomol.LevelFatal, nil, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(gomol.LevelInfo, nil, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(gomol.LevelWarning, nil, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(gomol.LevelError, nil, format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(gomol.LevelFatal, nil, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) PrintWithFields(fields lfi.Fields, format string, args ...interface{}) {}

func (l *Logger) PanicWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) DebugWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelDebug, fields, format, args...)
}

func (l *Logger) InfoWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelInfo, fields, format, args...)
}

func (l *Logger) WarnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelWarning, fields, format, args...)
}

func (l *Logger) ErrorWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelError, fields, format, args...)
}

func (l *Logger) FatalWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}
