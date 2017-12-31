package logger

import (
	"errors"
	"strings"
)

const (
	ConsoleTimeFormat = "2006-01-02 15:04:05.000"      // ConsoleTimeFormat
	JSONTimeFormat    = "2006-01-02T15:04:05.000-0700" // JSONTimeFormat
	// LOGRUS            Backend = "logrus"                       // Logrus backend
	// ZAP               Backend = "zap"                          // Zap backend
	// GOMOL             Backend = "gomol"                        // Gomol backend
)

var (
	LegalEncodings []string = []string{"console", "text", "json", "yaml"}
	LegalBackends  []string = []string{"gomol", "logrus", "zap"}
	LegalLevels    []string = []string{"debug", "info", "warning", "error", "fatal"}
)

var (
	ErrIllegalBackend  = errors.New("illegal log backend")
	ErrIllegalLevel    = errors.New("illegal log level")
	ErrIllegalEncoding = errors.New("illegal log encoding")
)

type Config struct {
	Backend       string   `env:"LOG_BACKEND" default:"gomol" yaml:"backend" json:"backend" toml:"backend"`
	Level         string   `env:"LOG_LEVEL" default:"info" yaml:"level" json:"level" toml:"level"`
	Encoding      string   `env:"LOG_ENCODING" default:"console" yaml:"encoding" json:"encoding" toml:"encoding"`
	DisableCaller bool     `env:"LOG_DISABLE_CALLER" yaml:"disable_caller" json:"disable_caller" toml:"disable_caller"`
	InitialFields Fields   `env:"LOG_FIELDS" yaml:"fields" json:"fields" toml:"fields"`
	Filters       []string `env:"LOG_FILTERS" yaml:"filters" json:"filters" toml:"filters"`
	OutputFile    string   `env:"LOG_OUTPUT_FILE" default:"./shared/logs/activity.log" yaml:"output_file" json:"output_file" toml:"output_file"`
}

func (c *Config) PostLoad() error {
	c.Level = strings.ToLower(c.Level)
	if !isLegalBackend(c.Backend) {
		return ErrIllegalBackend
	}
	if !isLegalLevel(c.Level) {
		return ErrIllegalLevel
	}
	if !isLegalEncoding(c.Encoding) {
		return ErrIllegalEncoding
	}
	return nil
}

func isLegalBackend(backend string) bool {
	for _, whitelisted := range LegalBackends {
		if backend == whitelisted {
			return true
		}
	}
	return false
}

func isLegalLevel(level string) bool {
	for _, whitelisted := range LegalLevels {
		if level == whitelisted {
			return true
		}
	}
	return false
}

func isLegalEncoding(encoding string) bool {
	for _, v := range LegalEncodings {
		if encoding == v {
			return true
		}
	}
	return false
}
