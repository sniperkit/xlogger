package core

var (
	LegalEncodings []string = []string{"console", "text", "json", "yaml"}
	LegalBackends  []string = []string{"logrus", "logrus-mate", "zap", "gomol", "seelog", "log4go", "glog", "go-logging", "log"}
	LegalLevels    []string = []string{"debug", "info", "warning", "error", "fatal"}
	LegalPlugins   []string = []string{"stats", "bench", "forward"}
)
