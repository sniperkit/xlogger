package config

import (
	"strings"

	lc "github.com/sniperkit/xlogger/pkg/core"
	lf "github.com/sniperkit/xlogger/pkg/fields"
)

type Config struct {
	// backends selection
	Backend     string `env:"LOG_BACKEND" default:"gomol" yaml:"backend" json:"backend" toml:"backend"`
	UseFallback bool   `env:"LOG_USE_FALLBACK" yaml:"use_fallback" json:"use_fallback" toml:"use_fallback"`
	// logger default config
	Level         string    `env:"LOG_LEVEL" default:"info" yaml:"level" json:"level" toml:"level"`
	Encoding      string    `env:"LOG_ENCODING" default:"console" yaml:"encoding" json:"encoding" toml:"encoding"`
	DisableCaller bool      `env:"LOG_DISABLE_CALLER" yaml:"disable_caller" json:"disable_caller" toml:"disable_caller"`
	InitialFields lf.Fields `env:"LOG_FIELDS" yaml:"fields" json:"fields" toml:"fields"`
	// output filters
	OutputFilters []string `env:"LOG_FILTERS" yaml:"log_filters" json:"log_filters" toml:"log_filters"`
	// write logs
	OutputFile       string `env:"LOG_OUTPUT_FILE" default:"./shared/logs/activity.log" yaml:"output_file" json:"output_file" toml:"output_file"`
	OutputFilePrefix string `env:"LOG_OUTPUT_FILE_PREFIX" default:"" yaml:"output_file_prefix" json:"output_file_prefix" toml:"output_file_prefix"`
	OutputFileSuffix string `env:"LOG_OUTPUT_FILE_SUFFIX" default:"" yaml:"output_file_suffix" json:"output_file_suffix" toml:"output_file_suffix"`
	OutputPrefixPath string `env:"LOG_OUTPUT_FILE" default:"./shared/logs/" yaml:"output_file" json:"output_prefix_path" toml:"output_prefix_path"`
	// watch dogs
	WatchKeywords []string `env:"LOG_WATCH_KEYWORDS" yaml:"log_watch_keywords" json:"log_watch_keywords" toml:"log_watch_keywords"`
	WatchPatterns []string `env:"LOG_WATCH_PATTERNS" yaml:"log_watch_patterns" json:"log_watch_patterns" toml:"log_watch_patterns"`
}

func (c *Config) PostLoad() error {
	c.Level = strings.ToLower(c.Level)
	if !isLegalBackend(c.Backend) {
		return lc.ErrIllegalBackend
	}
	if !isLegalLevel(c.Level) {
		return lc.ErrIllegalLevel
	}
	if !isLegalEncoding(c.Encoding) {
		return lc.ErrIllegalEncoding
	}
	return nil
}

func isLegalBackend(backend string) bool {
	for _, whitelisted := range lc.LegalBackends {
		if backend == whitelisted {
			return true
		}
	}
	return false
}

func isLegalLevel(level string) bool {
	for _, whitelisted := range lc.LegalLevels {
		if level == whitelisted {
			return true
		}
	}
	return false
}

func isLegalEncoding(encoding string) bool {
	for _, v := range lc.LegalEncodings {
		if encoding == v {
			return true
		}
	}
	return false
}
