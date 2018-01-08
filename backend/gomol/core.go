package gomol

import (
	"os"

	"github.com/aphistic/gomol"
	console "github.com/aphistic/gomol-console"

	log "github.com/sniperkit/xlogger/pkg"
	lcf "github.com/sniperkit/xlogger/pkg/config"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

const LoggerBackend = "gomol"

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

func New(c *lcf.Config) (*Logger, error) { // (*logger.Logger, error) {

	if c.Backend == "" {
		c.Backend = LoggerBackend
	}

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

func (Logger) Name() string {
	return LoggerBackend
}

func (l *Logger) WithFields(fields lfi.Fields) *Logger { // *logx.Logger {
	if len(fields) == 0 {
		return l
	}
	return &Logger{
		logger:        gomol.NewLogAdapterFor(l.logger, gomol.NewAttrsFromMap(fields)),
		disableCaller: l.disableCaller,
	}
}

func (l *Logger) Sync() error {
	return gomol.ShutdownLoggers()
}

func (l *Logger) log(level gomol.LogLevel, fields lfi.Fields, format string, args ...interface{}) {
	if fields == nil {
		fields = map[string]interface{}{}
	}
	if !l.disableCaller {
		fields["caller"] = log.GetCaller()
	}
	l.logger.Log(level, gomol.NewAttrsFromMap(fields.NormalizeTimeValues()), format, args...)
}
