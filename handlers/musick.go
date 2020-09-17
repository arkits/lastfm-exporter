package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arkits/musick/domain"
)

// NowPlayingController handles returning data when the /now endpoint is called
func NowPlayingController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&domain.LastFmPollingData)
}
