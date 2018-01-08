package gomol

import (
	"os"

	"github.com/aphistic/gomol"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

// not working properly yet
func (l *Logger) Println(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Panicln(format string, args ...interface{}) {
	l.log(gomol.LevelFatal, nil, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) Debugln(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Infoln(format string, args ...interface{}) {
	l.log(gomol.LevelInfo, nil, format, args...)
}

func (l *Logger) Warnln(format string, args ...interface{}) {
	l.log(gomol.LevelWarning, nil, format, args...)
}

func (l *Logger) Errorln(format string, args ...interface{}) {
	l.log(gomol.LevelError, nil, format, args...)
}

func (l *Logger) Fatalln(format string, args ...interface{}) {
	l.log(gomol.LevelFatal, nil, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) PrintlnWithFields(fields lfi.Fields, format string, args ...interface{}) {}

func (l *Logger) PaniclnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) DebuglnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelDebug, fields, format, args...)
}

func (l *Logger) InfolnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelInfo, fields, format, args...)
}

func (l *Logger) WarnlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelWarning, fields, format, args...)
}

func (l *Logger) ErrorlnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelError, fields, format, args...)
}

func (l *Logger) FatallnWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}
