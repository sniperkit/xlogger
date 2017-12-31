package log4go

import (
	// "time"

	"github.com/alecthomas/log4go"
	"github.com/sniperkit/logger"
)

// To finish
type Logger struct {
	logger        *log4go.Logger
	disableCaller bool
}

func New(c *logger.Config) (*Logger, error) {

	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	backend := seelog.log
	backend.Level = level

	switch c.Encoding {
	case "console":
		backend.Formatter = seelog.NewFormatter()

	case "json":
		backend.Formatter = seelog.NewFormatter()

	}
	factory := &Logger{
		logger:        backend.WithFields(nil),
		disableCaller: c.DisableCaller,
	}

	return factory.WithFields(c.InitialFields), nil
}

func (l *Logger) WithFields(fields logger.Fields) *Logger {
	if len(fields) == 0 {
		return l
	}
	backend := &Logger{
		logger: l.getLogger(fields),
	}
	return backend
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) DebugWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Debugf(format, args...)
}

func (l *Logger) InfoWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Infof(format, args...)
}

func (l *Logger) WarningWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Warnf(format, args...)
}

func (l *Logger) ErrorWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Errorf(format, args...)
}

func (l *Logger) FatalWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Fatalf(format, args...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}
