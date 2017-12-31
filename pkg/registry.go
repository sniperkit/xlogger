package logger

import (
	"fmt"
	"reflect"
	"strings"

	lcf "github.com/sniperkit/logger/pkg/config"
)

// Initialize creates a Logger Store object, initializing the logger
type Initialize func(addrs []string, options *lcf.Config) (Logger, error)

type Backend struct {
	Logger Logger
}

// NewBackend returns a new Backend
func NewBackend(l Logger) *Backend {
	return &Backend{Logger: l}
}

func registerLogger(logger Logger) {
	loggers[Name(logger)] = logger
}

// Name returns the name of a logger
func Name(logger Logger) string {
	parts := strings.Split(reflect.TypeOf(logger).String(), ".")
	return strings.ToLower(parts[len(parts)-1])
}

// ForName returns the service for a given name, or an error if it doesn't exist
func ForName(name string) (Logger, error) {
	if logger, ok := loggers[strings.ToLower(name)]; ok {
		return logger, nil
	}
	return &NotFound{}, fmt.Errorf("Logger '%s' not found", name)
}
