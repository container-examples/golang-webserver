package main

import (
	"flag"

	"github.com/container-examples/golang-webserver/config"
)

// ParseCommandLine parse flags and args from cli.
func ParseCommandLine() *config.Config {
	cfg := &config.Config{}

	// Config file flag
	flag.StringVar(&cfg.ConfigFile, "config.file", "", "Promxy configuration file path.")
	// Web & Log Flags
	flag.StringVar(&cfg.Web.ListenAddress, "web.listen-address", config.DefaultConfig.Web.ListenAddress, "Address to listen webserver.")
	flag.DurationVar(&cfg.Web.Timeout, "web.timeout", config.DefaultConfig.Web.Timeout, "Maximum duration before timing out requests.")
	flag.StringVar(&cfg.Web.Log.Level, "log.level", config.DefaultConfig.Web.Log.Level, "log level flags allowed [trace, debug, info, warn, error, fatal]")
	flag.StringVar(&cfg.Web.Log.Format, "log.format", config.DefaultConfig.Web.Log.Format, "log format flags allowed [logfmt, json]")

	// Parse flag parameters
	flag.Parse()

	return cfg
}
