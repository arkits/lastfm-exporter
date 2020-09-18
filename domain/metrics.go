package domain

import "github.com/prometheus/client_golang/prometheus"

var (
	// LastFmPollCounter counts the number of time the LastFM API was polled
	LastFmPollCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "musick_lastfm_polling_count",
			Help: "Musick - Number of times the LastFM API was polled",
		},
	)

	HttpRequestDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "http_request_duration_seconds",
			Buckets: []float64{0.1, 0.5, 1, 1.5, 2.0},
		},
		[]string{"path", "route"},
	)

	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
		},
		[]string{"path", "route"},
	)
)

func init() {
	prometheus.MustRegister(LastFmPollCounter)
	prometheus.MustRegister(HttpRequestDurations)
	prometheus.MustRegister(HttpRequestsTotal)

	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}
