package main

import (
	"time"

	"github.com/sniperkit/xlogger/backend/gomol"
	"github.com/sniperkit/xlogger/backend/logrus"
	"github.com/sniperkit/xlogger/backend/zap"

	logger "github.com/sniperkit/xlogger/pkg"
	logger_cfg "github.com/sniperkit/xlogger/pkg/config"
	logger_fields "github.com/sniperkit/xlogger/pkg/fields"
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
	logxFields logger_fields.Fields
	logxConfig *logger_cfg.Config
)

func main() {

	logxFields = logger.Fields{"appname": "multi-logger"}

	logxConfig = &logger_cfg.Config{
		Level:         logxLevel,
		Encoding:      logxEncoding,
		DisableCaller: logxDisabledCaller,
		OutputFile:    logxOutputFile,
		InitialFields: logxFields,
	}

	var err error
	switch logxEngine {
	case "gomol":
		logx, err = gomol.New(logxConfig)
	case "logrus":
		logx, err = logrus.New(logxConfig)
	case "zap":
		logx, err = zap.New(logxConfig)
	}
	if err != nil {
		panic(err)
	}

	logx.Info("test logger")
	logx.WarningWithFields(logger_fields.Fields{"now": time.Now(), "event": "event-1"}, "test logger")

}
