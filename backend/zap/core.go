package zap

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	lcf "github.com/sniperkit/xlogger/pkg/config"
	lco "github.com/sniperkit/xlogger/pkg/core"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

const LoggerBackend = "zap"

type Logger struct {
	logger *zap.SugaredLogger
}

func New(c *lcf.Config) (*Logger, error) {

	if c.Backend == "" {
		c.Backend = LoggerBackend
	}

	var (
		level        zap.AtomicLevel
		levelEncoder zapcore.LevelEncoder
		timeEncoder  zapcore.TimeEncoder
	)
	if err := level.UnmarshalText([]byte(c.Level)); err != nil {
		return nil, err
	}

	switch c.Encoding {
	case "console":
		levelEncoder = zapcore.CapitalColorLevelEncoder
		timeEncoder = zapConsoleTimeEncoder
	case "yaml":
		fallthrough

	case "json":
		fallthrough

	default:
		levelEncoder = zapcore.LowercaseLevelEncoder
		timeEncoder = zapJSONTimeEncoder
	}

	config := zap.Config{
		Level:             level,
		DisableCaller:     c.DisableCaller,
		Encoding:          c.Encoding,
		Development:       false,
		DisableStacktrace: true,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			MessageKey:     "message",
			CallerKey:      "caller",
			EncodeLevel:    levelEncoder,
			EncodeTime:     timeEncoder,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	factory, err := config.Build()
	if err != nil {
		return nil, err
	}
	sugaredLogger := factory.WithOptions(zap.AddCallerSkip(1)).Sugar()

	backend := (&Logger{logger: sugaredLogger}).WithFields(c.InitialFields)
	return backend, nil
}

func (Logger) Name() string {
	return LoggerBackend
}

func (l *Logger) WithFields(fields lfi.Fields) *Logger {
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

func (l *Logger) DebugWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Debugf(format, args...)
}

func (l *Logger) InfoWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Infof(format, args...)
}

func (l *Logger) WarningWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Warnf(format, args...)
}

func (l *Logger) ErrorWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Errorf(format, args...)
}

func (l *Logger) FatalWithFields(fields lfi.Fields, format string, args ...interface{}) {
	l.getLogger(fields).Fatalf(format, args...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) getLogger(fields lfi.Fields) *zap.SugaredLogger {
	if len(fields) == 0 {
		return l.logger
	}
	return l.logger.With(flatten(fields)...)
}

func flatten(fields lfi.Fields) []interface{} {
	flattened := []interface{}{}
	for key, value := range fields {
		flattened = append(flattened, key)
		flattened = append(flattened, value)
	}
	return flattened
}

func zapConsoleTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(lco.ConsoleTimeFormat))
}

func zapJSONTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(lco.JSONTimeFormat))
}
