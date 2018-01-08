package gomol

import (
	"os"

	"github.com/aphistic/gomol"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

// not working properly yet

func (l *Logger) Printf(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log(gomol.LevelFatal, nil, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(gomol.LevelInfo, nil, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(gomol.LevelWarning, nil, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(gomol.LevelError, nil, format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(gomol.LevelFatal, nil, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) PrintfWithFields(fields lfi.Fields, format string, args ...interface{}) {}

func (l *Logger) PanicfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) DebugfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelDebug, fields, format, args...)
}

func (l *Logger) InfofWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelInfo, fields, format, args...)
}

func (l *Logger) WarnfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelWarning, fields, format, args...)
}

func (l *Logger) ErrorfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelError, fields, format, args...)
}

func (l *Logger) FatalfWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}
