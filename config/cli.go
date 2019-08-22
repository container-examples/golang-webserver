package config

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

var (
	levels = regexp.MustCompile("^(debug|info|warn|error|fatal)$")
	json   = regexp.MustCompile("^json$")
)

// LogFlagParse level logs
func (c *Config) LogFlagParse() *logrus.Logger {
	logger := logrus.New()

	if levels.MatchString(c.Web.Log.Level) {
		lvl, err := logrus.ParseLevel(c.Web.Log.Level)
		if err != nil {
			// Should never happen since we select correct levels
			// Unless logrus commits a breaking change on level names
			logger.Fatalf("Invalid log level: %s", err.Error())
		}

		logger.SetLevel(lvl)
		logger.WithField("level", lvl).Debugln("Set level logging")
	} else {
		logger.WithField("level", c.Web.Log.Level).Warnln("Log level setted doesn't exist (default: Info)")
	}

	if json.MatchString(c.Web.Log.Format) {
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.WithField("format", "json").Debugln("Set format logging")
	}

	return logger
}
