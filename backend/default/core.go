package log

import (
	"errors"
	"log"

	lcf "github.com/sniperkit/xlogger/pkg/config"
	lfi "github.com/sniperkit/xlogger/pkg/fields"
)

const (
	LoggerBackend              = "log"
	LoggerDefaultPrefix string = ""
	LoggerDefaultFlag   int    = 1
)

type Logger struct {
	logger *log.Logger
}

func New(c *lcf.Config) (*Logger, error) {
	if c.Backend == "" {
		c.Backend = LoggerBackend
	}
	backend := (&Logger{logger: log.New(nil, LoggerDefaultPrefix, LoggerDefaultFlag)}).WithFields(c.InitialFields)
	return backend, nil
}

func (Logger) Name() string { return LoggerBackend }

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (l *Logger) WithFields(fields lfi.Fields) *Logger { return l }
func (l *Logger) Sync() error                          { return errors.New("Not implemented") }

func flatten(fields lfi.Fields) []interface{} {
	flattened := []interface{}{}
	for key, value := range fields {
		flattened = append(flattened, key)
		flattened = append(flattened, value)
	}
	return flattened
}
