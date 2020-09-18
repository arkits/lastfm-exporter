package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arkits/musick/domain"
	"github.com/arkits/musick/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

func init() {

	// Setup the Application wide config through Viper
	SetupConfig()

	// Kick off PollRecentTracks
	go domain.PollRecentTracks()
}

func main() {

	port := viper.GetString("server.port")
	serviceName := viper.GetString("server.name")

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.VersionController).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/%s", serviceName), handlers.VersionController).Methods(http.MethodGet)

	r.HandleFunc(fmt.Sprintf("/%s/now", serviceName), handlers.NowPlayingController).Methods(http.MethodGet)

	// Expose Prometheus metrics
	r.HandleFunc(fmt.Sprintf("/%s/metrics", serviceName), promhttp.Handler().ServeHTTP).Methods(http.MethodGet)

	r.Use(handlers.LoggingMiddleware)
	r.Use(handlers.MetricsMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))

	log.Printf("Starting musick on http://localhost:%v", port)
	http.ListenAndServe(":"+port, r)

}

// SetupConfig -  Setup the application config by reading the config file via Viper
func SetupConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file! - %s", err)
	}

}
