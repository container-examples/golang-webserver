package webserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (h *Handler) Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			start, end time.Time
			lrw        *loggingResponseWriter
		)

		start = time.Now().UTC()
		lrw = NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		end = time.Now().UTC()
		latency := end.Sub(start)
		h.Logger.WithFields(logrus.Fields{
			"status":    lrw.statusCode,
			"method":    r.Method,
			"path":      r.URL.Path,
			"addr":      r.RemoteAddr,
			"duration":  latency.Seconds(),
			"userAgent": r.Header.Get("User-Agent"),
			"protocol":  r.Proto,
			"length":    r.ContentLength,
		}).Info("request")
	}
}
