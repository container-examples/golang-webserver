package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/imdario/mergo"
	"github.com/sirupsen/logrus"

	"github.com/container-examples/golang-webserver/config"
	"github.com/container-examples/golang-webserver/webserver"
)

var (
	logger      *logrus.Logger
	cliCfg, cfg *config.Config
	err         error
)

func reload(cliCfg *config.Config, logger *logrus.Logger) (*config.Config, error) {
	cfg := &config.DefaultConfig

	// Parse config file if needed
	if cliCfg.ConfigFile != "" {
		logger.Info("Loading configuration file")
		fileCfg, err := config.LoadFile(logger, cliCfg.ConfigFile)
		if err != nil {
			logger.Errorln("Error loading config file", err)
			return nil, err
		}

		cfg = fileCfg
	}

	// Merge overwritting cliCfg into cfg
	if err := mergo.MergeWithOverwrite(cliCfg, cfg); err != nil {
		logger.Errorln("Error merging config file with flags", err)
		return nil, err
	}

	return cfg, nil
}

func main() {
	var (
		h    *webserver.Handler
		done chan os.Signal
	)

	logger.Infoln("Start webserver ...")

	// Loading Config file & merge Cli args
	cfg, err = reload(cliCfg, logger)
	if err != nil {
		logger.Fatalln("Impossible to load configuration")
	}

	// Run HTTP server
	h = webserver.New(cfg, logger)
	h.Router.HandleFunc("/", h.Logging(h.Hello)).Methods("GET")

	// Graceful Notify
	done = make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.WithField("addr", cfg.Web.ListenAddress).Warnf("Starting Webserver")
		if err := h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	<-done
	logger.Warn("Stopping Webserver ...")

	// Shutdown HTTP Server
	h.Shutdown()
}

func init() {
	// Parse Flags
	cliCfg = ParseCommandLine()
	// Parse level & format logs
	logger = cliCfg.LogFlagParse()
}
