package webserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/container-examples/golang-webserver/config"
)

// Handler serves various HTTP endpoints of the remote adapter server
type Handler struct {
	Logger *logrus.Logger
	Server *http.Server
	Config *config.Config
	Router *mux.Router
}

func (h *Handler) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := h.Server.Shutdown(ctx); err != nil {
		h.Logger.Fatalf("Server Shutdown Failed:%+v", err)
	}

	h.Logger.Info("Server Exited Properly")
}

// New HTTP server
func New(cfg *config.Config, logger *logrus.Logger) *Handler {
	var router = mux.NewRouter()

	return &Handler{
		Logger: logger,
		Router: router,
		Config: cfg,
		Server: &http.Server{
			Addr: cfg.Web.ListenAddress,
			Handler: handlers.CORS(
				handlers.AllowedMethods([]string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodDelete,
				}),
				handlers.AllowedOrigins([]string{"*"}),
			)(router),
		},
	}
}
