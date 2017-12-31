package logger

import (
	log "github.com/sniperkit/xlogger/backend/default"
	lc "github.com/sniperkit/xlogger/pkg/config"
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
