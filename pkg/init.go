package logger

import (
	log "github.com/sniperkit/logger/backend/default"
	lc "github.com/sniperkit/logger/pkg/config"
)

var (
	logx    Logger
	loggers = make(map[string]Logger)
)

func init() {
	logx = &NotFound{}
}

func New(c *lc.Config) (Logger, error) {
	return log.New(&lc.Config{
		Encoding:      "console",
		InitialFields: nil,
		DisableCaller: true,
	})
}
