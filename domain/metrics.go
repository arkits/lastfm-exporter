package domain

import "github.com/prometheus/client_golang/prometheus"

var (
	// HTTPRequestDurations - Duration of HTTP requests in seconds
	HTTPRequestDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: []float64{0.1, 0.5, 1, 1.5, 2.0},
		},
		[]string{"path", "route"},
	)

	// HTTPRequestTotal - Counter for total requests received
	HTTPRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Counter for total requests received",
		},
		[]string{"path", "route"},
	)

	// LastFmPollCounter - Number of times the LastFM API was polled
	LastFmPollCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "musick_lastfm_polling_count",
			Help: "Musick - Number of times the LastFM API was polled",
		},
	)

	// LastFmPollErrorCounter - Number of times the LastFM API polling threw an error
	LastFmPollErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "musick_lastfm_polling_error_count",
			Help: "Musick - Number of times the LastFM API polling threw an error",
		},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequestDurations)
	prometheus.MustRegister(HTTPRequestTotal)

	prometheus.MustRegister(LastFmPollCounter)
	prometheus.MustRegister(LastFmPollErrorCounter)

	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}
