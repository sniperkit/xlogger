package main

import (
	"time"

	logger "github.com/sniperkit/logger/pkg"
	logger_cfg "github.com/sniperkit/logger/pkg/config"
	logger_fields "github.com/sniperkit/logger/pkg/fields"

	"github.com/sniperkit/logger/backend/gomol"
)

const (
	logxLevel          string = "info"
	logxEncoding       string = "console" // json
	logxDisabledCaller bool   = false
	logxOutputFile     string = "./logs/activity.log"
)

var (
	logx       logger.Logger
	logxConfig *logger_cfg.Config = &logger_cfg.Config{}
)

func main() {

	logxConfig.Level = logxLevel
	logxConfig.Encoding = logxEncoding
	logxConfig.DisableCaller = logxDisabledCaller
	logxConfig.OutputFile = logxOutputFile
	logxConfig.InitialFields = logger_fields.Fields{"appname": "logger-gomol"}

	var err error
	logx, err = gomol.New(logxConfig)
	if err != nil {
		panic(err)
	}

	logx.Info("test logger")
	logx.WarningWithFields(logger_fields.Fields{"now": time.Now(), "event": "event-1"}, "test logger")

}
