package config

import (
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// DefaultConfig if not file is indicated
var DefaultConfig = Config{
	Web: webOptions{
		ListenAddress: "0.0.0.0:3000",
		Timeout:       5 * time.Second,
		Log: logOptions{
			Level:  "info",
			Format: "logfmt",
		},
	},
}

// LoadFile parses the given YAML file into a Config.
func LoadFile(logger *logrus.Logger, filename string) (*Config, error) {
	cfg := &Config{}
	*cfg = DefaultConfig

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(content, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
