package logger

import (
	"fmt"
	"reflect"
	"strings"
	// "sort"

	"github.com/fatih/color"
	"github.com/hoop33/entrevista"
	// "github.com/sniperkit/logger/backends/zap"
)

// Initialize creates a Logger Store object, initializing the logger
type Initialize func(addrs []string, options *Config) (Logger, error)

var (
	logx    Logger
	loggers = make(map[string]Logger)
)

type Backend struct {
	Logger Logger
}

// NewBackend returns a new Backend
func NewBackend(l Logger) *Backend {
	return &Backend{Logger: l}
}

/*
func init() {
	logx = zap.New(nil)
}
*/

/*
	// Backend initializers
	initializers = make(map[Backend]Initialize)

	supportedBackend = func() string {
		keys := make([]string, 0, len(initializers))
		for k := range initializers {
			keys = append(keys, string(k))
		}
		sort.Strings(keys)
		return strings.Join(keys, ", ")
	}()
)
*/

/*
// NewLogger creates an instance of logger
func NewLogger(backend Backend, addrs []string, options *Config) (Logger, error) {
	if init, exists := initializers[backend]; exists {
		return init(addrs, options)
	}
	return nil, fmt.Errorf("%s %s", ErrIllegalBackend.Error(), supportedBackend)
}

// AddStore adds a new store backend to libkv
func AddBackend(backend Backend, init Initialize) {
	initializers[backend] = init
}
*/

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

func createInterview() *entrevista.Interview {
	interview := entrevista.NewInterview()
	interview.ShowOutput = func(message string) {
		fmt.Print(color.GreenString(message))
	}
	interview.ShowError = func(message string) {
		color.Red(message)
	}
	return interview
}
