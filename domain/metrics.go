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
)

func init() {
	prometheus.MustRegister(LastFmPollCounter)
}
