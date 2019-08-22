package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

// logOptions struct parameters for logging level and format
type logOptions struct {
	Level  string `yaml:"level,omitempty"`
	Format string `yaml:"format,omitempty"`
}

type webOptions struct {
	ListenAddress string        `yaml:"listen_address,omitempty"`
	Timeout       time.Duration `yaml:"timeout,omitempty"`
	Log           logOptions    `yaml:"log,omitempty"`

	// Catches all undefined fields and must be empty after parsing.
	XXX map[string]interface{} `yaml:",inline"`
}

type Config struct {
	ConfigFile string
	Logger     *logrus.Logger
	Web        webOptions `yaml:"web,omitempty"`
}
