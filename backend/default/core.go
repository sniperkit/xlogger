package dlog

import (
	"errors"
	"log"

	lcf "github.com/sniperkit/logger/pkg/config"
	lfi "github.com/sniperkit/logger/pkg/fields"
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

func (l *Logger) WithFields(fields lfi.Fields) *Logger                                    { return l }
func (l *Logger) Debug(format string, args ...interface{})                                {}
func (l *Logger) Info(format string, args ...interface{})                                 {}
func (l *Logger) Warning(format string, args ...interface{})                              {}
func (l *Logger) Error(format string, args ...interface{})                                {}
func (l *Logger) Fatal(format string, args ...interface{})                                {}
func (l *Logger) DebugWithFields(fields lfi.Fields, format string, args ...interface{})   {}
func (l *Logger) InfoWithFields(fields lfi.Fields, format string, args ...interface{})    {}
func (l *Logger) WarningWithFields(fields lfi.Fields, format string, args ...interface{}) {}
func (l *Logger) ErrorWithFields(fields lfi.Fields, format string, args ...interface{})   {}
func (l *Logger) FatalWithFields(fields lfi.Fields, format string, args ...interface{})   {}
func (l *Logger) Sync() error                                                             { return errors.New("Not implemented") }

func flatten(fields lfi.Fields) []interface{} {
	flattened := []interface{}{}
	for key, value := range fields {
		flattened = append(flattened, key)
		flattened = append(flattened, value)
	}
	return flattened
}
