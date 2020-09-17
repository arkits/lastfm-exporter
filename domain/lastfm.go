package domain

import (
	"log"
	"sync"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
	"github.com/spf13/viper"
)

// PollingData encapsulates the data that needs to be persisted when polling a particular data source
type PollingData struct {
	mu          sync.RWMutex
	lastUpdated time.Time
	pollCount   int
	User        string `json:"username"`
	NowPlaying  struct {
		TrackName   string `json:"trackName"`
		ArtistName  string `json:"artistName"`
		AlbumName   string `json:"albumName"`
		LastFmURL   string `json:"lastFmUrl"`
		CoverArtURL string `json:"coverArtUrl"`
	} `json:"nowPlaying"`
}

// LastFmPollingData persists the Polling Data from LastFm
var LastFmPollingData PollingData

// PollRecentTracks - polls recent LastFM tracks
func PollRecentTracks() {

	lastFmAPI := lastfm.New(
		viper.GetString("lastfm.apiKey"),
		viper.GetString("lastfm.apiSecret"),
	)

	for {
		log.Printf("Getting recent tracks from LastFM - PollCount=%v", LastFmPollingData.pollCount)

		result, err := lastFmAPI.User.GetRecentTracks(
			lastfm.P{
				"user": viper.GetString("lastfm.username"),
			},
		)

		if err != nil {
			log.Fatal(err)
			break
		}

		LastFmPollingData.mu.Lock()

		// Persist the relevant data
		LastFmPollingData.User = result.User
		LastFmPollingData.NowPlaying.TrackName = result.Tracks[0].Name
		LastFmPollingData.NowPlaying.AlbumName = result.Tracks[0].Album.Name
		LastFmPollingData.NowPlaying.ArtistName = result.Tracks[0].Artist.Name
		LastFmPollingData.NowPlaying.LastFmURL = result.Tracks[0].Url
		LastFmPollingData.NowPlaying.CoverArtURL = result.Tracks[0].Images[len(result.Tracks[0].Images)-1].Url

		// update when the last update took place
		LastFmPollingData.lastUpdated = time.Now()

		// update the pollCount
		LastFmPollingData.pollCount++

		LastFmPollingData.mu.Unlock()

		time.Sleep(1 * time.Second)
	}
}
