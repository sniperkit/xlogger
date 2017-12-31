package logrus_mate

import (
	"github.com/sirupsen/logrus"
	"github.com/sniperkit/logrus_mate"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	lc "github.com/sniperkit/logger/pkg/config"
	lf "github.com/sniperkit/logger/pkg/fields"
)

const LoggerBackend = "logrus-mate"

type Logger struct {
	entry         *logrus.Entry
	disableCaller bool
	isFiltered    bool
}

func New(c *lc.Config) (*Logger, error) {

	if c.Backend == "" {
		c.Backend = LoggerBackend
	}

	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}

	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
	)

	backend := logrus.New()

	mate.Hijack(
		backend,
		"mike",
	)

	backend.Level = level

	switch c.Encoding {
	case "console":
		formatter := &prefixed.TextFormatter{
			FullTimestamp:    true,
			TimestampFormat:  logger.ConsoleTimeFormat,
			QuoteEmptyFields: true,
		}
		backend.Formatter = formatter

	case "json":
		backend.Formatter = &logrus.JSONFormatter{
			TimestampFormat: logger.JSONTimeFormat,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
			},
		}
	}

	factory := &Logger{
		entry:         backend.WithFields(nil),
		disableCaller: c.DisableCaller,
		isFiltered:    len(fieldsFilter) > 0,
		// isFiltered:    len(c.Filters) > 0,
	}
	return factory.WithFields(c.InitialFields), nil
}

func (Logger) Name() string {
	return LoggerBackend
}

func (l *Logger) WithFields(fields lf.Fields) *Logger {
	if len(fields) == 0 {
		return l
	}
	return &Logger{
		entry:         l.getEntry(fields),
		disableCaller: l.disableCaller,
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.decorateEntry(nil).Debugf(format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.decorateEntry(nil).Infof(format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.decorateEntry(nil).Warningf(format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.decorateEntry(nil).Errorf(format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.decorateEntry(nil).Fatalf(format, args...)
}

func (l *Logger) DebugWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Debugf(format, args...)
}

func (l *Logger) InfoWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Infof(format, args...)
}

func (l *Logger) WarningWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Warningf(format, args...)
}

func (l *Logger) ErrorWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Errorf(format, args...)
}

func (l *Logger) FatalWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.decorateEntry(fields).Fatalf(format, args...)
}

func (l *Logger) Sync() error {
	return nil
}

func (l *Logger) getEntry(fields logger.Fields) *logrus.Entry {
	if len(fields) == 0 {
		return l.entry
	}
	return l.entry.WithFields(logrus.Fields(fields.NormalizeTimeValues()))
}

func (l *Logger) decorateEntry(fields logger.Fields) *logrus.Entry {
	entry := l.getEntry(fields)
	if l.disableCaller {
		return entry
	}
	return entry.WithFields(logrus.Fields(map[string]interface{}{
		"caller": logger.GetCaller(),
	}))
}
