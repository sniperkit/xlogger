package gomol

import (
	"os"

	"github.com/aphistic/gomol"
	console "github.com/aphistic/gomol-console"

	"github.com/sniperkit/logger"
)

type Logger struct {
	logger        gomol.WrappableLogger
	disableCaller bool
}

/*
// Register loggers gomol to logger
func Register() {
	logger.AddBackend(logger.GOMOL, New)
}
*/

func New(c *logger.Config) (*Logger, error) { // (*logger.Logger, error) {
	level, _ := gomol.ToLogLevel(c.Level)
	gomol.SetLogLevel(level)

	switch c.Encoding {
	case "console":
		consoleLogger, err := console.NewConsoleLogger(
			&console.ConsoleLoggerConfig{
				Colorize: true,
				Writer:   os.Stderr,
			})
		if err != nil {
			return nil, err
		}
		tpl, err := gomol.NewTemplate("" +
			"[{{.Timestamp.Format \"2006-01-02 15:04:05.000\"}}] " +
			`{{color}}{{printf "%5s" (ucase .LevelName)}}{{reset}} ` +
			"{{.Message}}" +
			"{{if .Attrs}}{{range $key, $val := .Attrs}} {{color}}{{$key}}{{reset}}={{$val}}{{end}}{{end}}")
		if err != nil {
			return nil, err
		}
		consoleLogger.SetTemplate(tpl)
		gomol.AddLogger(consoleLogger)

	case "yaml":
		fallthrough

	case "json":
		fallthrough

	default:
		gomol.AddLogger(newJSONLogger())

	}
	if err := gomol.InitLoggers(); err != nil {
		return nil, err
	}
	factory := &Logger{
		logger:        gomol.NewLogAdapter(nil),
		disableCaller: c.DisableCaller,
	}

	return factory.WithFields(c.InitialFields), nil
}

func (l *Logger) WithFields(fields logger.Fields) *Logger { // *logger.Logger {
	if len(fields) == 0 {
		return l
	}
	return &Logger{
		logger:        gomol.NewLogAdapterFor(l.logger, gomol.NewAttrsFromMap(fields)),
		disableCaller: l.disableCaller,
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(gomol.LevelDebug, nil, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(gomol.LevelInfo, nil, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
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

func (l *Logger) DebugWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelDebug, fields, format, args...)
}

func (l *Logger) InfoWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelInfo, fields, format, args...)
}

func (l *Logger) WarningWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelWarning, fields, format, args...)
}

func (l *Logger) ErrorWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelError, fields, format, args...)
}

func (l *Logger) FatalWithFields(fields logger.Fields, format string, args ...interface{}) {
	l.log(gomol.LevelFatal, fields, format, args...)
	l.logger.ShutdownLoggers()
	os.Exit(1)
}

func (l *Logger) Sync() error {
	return gomol.ShutdownLoggers()
}

func (l *Logger) log(level gomol.LogLevel, fields logger.Fields, format string, args ...interface{}) {
	if fields == nil {
		fields = map[string]interface{}{}
	}
	if !l.disableCaller {
		fields["caller"] = logger.GetCaller()
	}
	l.logger.Log(level, gomol.NewAttrsFromMap(fields.NormalizeTimeValues()), format, args...)
}