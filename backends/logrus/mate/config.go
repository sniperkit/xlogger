package logrus_mate

const (
	DefaultHijackerName         string = "rosco"
	DefaultHijackerConfigString        = `{ mike {formatter.name = "json"} }`
	DefaultHijackerConfigFile   string = "mate.conf"
)

var (
	LegalMateHijackModes map[string]string = map[string]string{
		"SL":   "Hijack logrus.StandardLogger()",
		"NL":   "Create new logger from mate",
		"LBM":  "Hijack logger by mate",
		"FCS":  "Fallback the ConfigString",
		"FCWH": "Fallback config while hijack standard logger",
	}
	LegalMateHooks map[string][]string = map[string][]string{
		"airbrake":  []string{"project-id", "api-key", "env"},
		"syslog":    []string{"network", "address", "priority", "tag"},
		"bugsnag":   []string{"api-key"},
		"slackrus":  []string{"url", "levels", "channel", "emoji", "username"},
		"graylog":   []string{"address", "facility", "extra"},
		"mail":      []string{"app-name", "host", "port", "from", "to", "username", "password"},
		"beat":      []string{"app-name", "protocol", "address", "always-sent-fields", "prefix"}, // to check
		"logstash":  []string{"app-name", "protocol", "address", "always-sent-fields", "prefix"},
		"file":      []string{"filename", "max-lines", "max-size", "daily", "max-days", "rotate", "level"},
		"bearychat": []string{"url", "levels", "channel", "user", "markdown", "async"},
	}
	LegalMateFormatters map[string][]string = map[string][]string{
		"text": []string{"force-colors", "disable-colors", "disable-timestamp", "full-timestamp", "timestamp-format", "disable-sorting"},
		"json": []string{"timestamp-format"},
		// "yaml": []string{"timestamp-format"},
	}
	LegalMateFormattersAliases map[string]string = map[string]string{
		"text": "console",
	}
	LegalMateWriters map[string][]string = map[string][]string{
		"redisio": []string{"network", "password", "db", "list-name", "address", "timeout"},
		// "elastic":  []string{"network", "token", "user", "password", "db", "address", "timeout"},
		// "mongodb":  []string{"network", "user", "password", "db", "address", "timeout"},
		// "mysql":    []string{"network", "user", "password", "db", "address", "timeout"},
		// "postgres": []string{"network", "user", "password", "db", "address", "timeout"},
		// "mssql":    []string{"network", "user", "password", "db", "address", "timeout"},
	}
)
