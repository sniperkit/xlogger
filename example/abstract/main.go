package main

import (
	"time"

	logger "github.com/sniperkit/logger/pkg"
	logger_cfg "github.com/sniperkit/logger/pkg/config"
	logger_fields "github.com/sniperkit/logger/pkg/fields"
)

const (
	logxEngine         string = "zap"
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
		Backend:       logxEngine, // specifiy the logging backend
		Level:         logxLevel,
		Encoding:      logxEncoding,
		DisableCaller: logxDisabledCaller,
		OutputFile:    logxOutputFile,
		InitialFields: logger_fields.Fields{"appname": "logger-abstract"},
	}

	var err error
	logx, err = logger.New(logxConfig)
	if err != nil {
		panic(err)
	}

	logx.Info("test logger abstract")
	logx.WarningWithFields(logger_fields.Fields{"now": time.Now(), "event": "event-1"}, "test logger")

}
