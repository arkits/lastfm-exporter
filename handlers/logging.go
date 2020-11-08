package handlers

import (
	"net/http"
	"time"

	"github.com/arkits/musick/domain"
)

// LoggingMiddleware logs incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("%s - %s | %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// MetricsMiddleware logs metrics related to the HTTP requests and responses
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func(begin time.Time) {
			domain.HTTPRequestDurations.WithLabelValues(r.Method, r.RequestURI).Observe(time.Since(begin).Seconds())
			domain.HTTPRequestTotal.WithLabelValues(r.Method, r.RequestURI).Inc()
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}
