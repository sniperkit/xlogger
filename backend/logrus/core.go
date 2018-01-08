package logrus

import (
	"os"

	"github.com/rifflock/lfshook"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	// "github.com/sniperkit/xlogger/backends/logrus/filtered"
	// "github.com/sniperkit/logrus_mate"
	// _ "github.com/sniperkit/logrus_mate/hooks/file"

	lp "github.com/sniperkit/xlogger/pkg"
	lcf "github.com/sniperkit/xlogger/pkg/config"
	lco "github.com/sniperkit/xlogger/pkg/core"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

const LoggerBackend = "logrus"

type Logger struct {
	entry         *logrus.Entry
	disableCaller bool
	isFiltered    bool
}

func New(c *lcf.Config) (*Logger, error) {

	if c.Backend == "" {
		c.Backend = LoggerBackend
	}

	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./shared/logs/info.log",
		logrus.WarnLevel:  "./shared/logs/warn.log",
		logrus.DebugLevel: "./shared/logs/debug.log",
		logrus.ErrorLevel: "./shared/logs/error.log",
	}

	backend := logrus.New()
	backend.Level = level
	// backend.Out = nil

	file, err := os.OpenFile("./shared/logs/combined.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		backend.Out = file
	} else {
		return nil, err
	}

	switch c.Encoding {
	case "console":
		formatter := &prefixed.TextFormatter{
			FullTimestamp:    true,
			TimestampFormat:  lco.ConsoleTimeFormat,
			QuoteEmptyFields: true,
		}

		backend.Formatter = formatter

	case "json":
		backend.Formatter = &logrus.JSONFormatter{
			TimestampFormat: lco.JSONTimeFormat,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
			},
		}
	}

	backend.Hooks.Add(lfshook.NewHook(
		pathMap,
		backend.Formatter,
	))

	/*
		// backend := logrus.New()
		mate, _ := logrus_mate.NewLogrusMate(
			logrus_mate.ConfigString(
				`{ rosco {formatter.name = "json"} }`,
			),
			logrus_mate.ConfigFile(
				"./shared/config/logrus-mate.conf", // { rosco {formatter.name = "text"} }
			),
		)

		mate.Hijack(
			backend,
			"rosco",
		)

		// backend.Errorln("hello std logger is hijack by rosco")
		// backend.Fatalln("hello std logger is hijack by rosco")

		// just for test
		fieldsFilter := []string{
			"password",
			"email",
		}

		if len(fieldsFilter) > 0 {
			//if len(c.Filters) > 0 {
			backend.Formatter = filtered.New(fieldsFilter, backend.Formatter)
		}
	*/

	factory := &Logger{
		entry:         backend.WithFields(nil),
		disableCaller: c.DisableCaller,
		// isFiltered:    len(fieldsFilter) > 0,
		// isFiltered:    len(c.Filters) > 0,
	}
	return factory.WithFields(c.InitialFields), nil
}

func (Logger) Name() string {
	return LoggerBackend
}

func (l *Logger) WithFields(fields lfi.Fields) *Logger {
	if len(fields) == 0 {
		return l
	}
	return &Logger{
		entry:         l.getEntry(fields),
		disableCaller: l.disableCaller,
	}
}

func (l *Logger) Sync() error {
	return nil
}

func (l *Logger) getEntry(fields lfi.Fields) *logrus.Entry {
	if len(fields) == 0 {
		return l.entry
	}
	return l.entry.WithFields(logrus.Fields(fields.NormalizeTimeValues()))
}

func (l *Logger) decorateEntry(fields lfi.Fields) *logrus.Entry {
	entry := l.getEntry(fields)
	if l.disableCaller {
		return entry
	}
	return entry.WithFields(logrus.Fields(map[string]interface{}{
		"caller": lp.GetCaller(),
	}))
}
