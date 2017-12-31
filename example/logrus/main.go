package main

import (
	"os"
	"time"

	logger "github.com/sniperkit/logger/pkg"
	logger_cfg "github.com/sniperkit/logger/pkg/config"
	logger_fields "github.com/sniperkit/logger/pkg/fields"

	"github.com/sniperkit/logger/backend/logrus"
)

const (
	logxLevel            string = "info"
	logxEncoding         string = "console" // console
	logxDisabledCaller   bool   = false
	logxOutputPrefixPath string = "./shared/logs"
	logxOutputFile       string = logxOutputPrefixPath + "./shared/logs/activity.log"
)

var (
	logx       logger.Logger
	logxConfig *logger_cfg.Config
)

func ensureDir(path string) {
	d, err := os.Open(path)
	if err != nil {
		os.MkdirAll(path, os.FileMode(0755))
	}
	d.Close()
}

func main() {

	ensureDir(logxOutputPrefixPath)
	logxConfig = &logger_cfg.Config{
		Level:         logxLevel,
		Encoding:      logxEncoding,
		DisableCaller: logxDisabledCaller,
		OutputFile:    logxOutputFile,
		InitialFields: logger_fields.Fields{"appname": "logger-logrus"},
	}

	var err error
	logx, err = logrus.New(logxConfig)
	if err != nil {
		panic(err)
	}

	logx.Info("test logger")
	logx.WarningWithFields(logger_fields.Fields{"now": time.Now(), "event": "event-1"}, "test logger")

}
