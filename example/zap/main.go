package main

import (
	"time"

	"github.com/sniperkit/xlogger/backend/zap"

	logger "github.com/sniperkit/xlogger/pkg"
	logger_cfg "github.com/sniperkit/xlogger/pkg/config"
	logger_fields "github.com/sniperkit/xlogger/pkg/fields"
)

const (
	logxLevel          string = "info"
	logxEncoding       string = "console" // json
	logxDisabledCaller bool   = false
	logxOutputFile     string = "./logs/activity.log"
)

var (
	logx       logger.Logger
	logxConfig *logger_cfg.Config
)

func main() {

	logxConfig = &logger_cfg.Config{
		Level:         logxLevel,
		Encoding:      logxEncoding,
		DisableCaller: logxDisabledCaller,
		OutputFile:    logxOutputFile,
		InitialFields: logger_fields.Fields{"appname": "logger-zap"},
	}

	var err error
	logx, err = zap.New(logxConfig)
	if err != nil {
		panic(err)
	}

	logx.Info("test logger")
	logx.WarningWithFields(logger_fields.Fields{"now": time.Now(), "event": "event-1"}, "test logger")

}
