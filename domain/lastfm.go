package domain

import (
	"log"
	"sync"
	"time"

	"git.maych.in/thunderbottom/lastfm-go"
	"git.maych.in/thunderbottom/lastfm-go/api/user"
	"github.com/spf13/viper"
)

// PollingData encapsulates the data that needs to be persisted when polling a particular data source
type PollingData struct {
	mu          sync.Mutex
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
	lastFmClient := lastfm.New(
		viper.GetString("lastfm.apiKey"),
		viper.GetString("lastfm.apiSecret"),
	)

	lastFmUser := user.New(&lastFmClient, viper.GetString("lastfm.username"))

	for {
		recentTracks, err := lastFmUser.GetRecentTracks(false, 1)

		if err != nil {
			log.Println(err)
			LastFmPollErrorCounter.Inc()
		} else {

			// The API can occasionally return a RecentTracks array of length 0...
			// Ignore and continue in this case
			if len(recentTracks.RecentTracks.Tracks) <= 0 {
				continue
			}

			lastTrack := recentTracks.RecentTracks.Tracks[0]

			// Begin updating LastFmPollingData
			LastFmPollingData.mu.Lock()

			// Persist the relevant data
			LastFmPollingData.User = viper.GetString("lastfm.username")
			LastFmPollingData.NowPlaying.TrackName = lastTrack.Name
			LastFmPollingData.NowPlaying.AlbumName = lastTrack.Album.Text
			LastFmPollingData.NowPlaying.ArtistName = lastTrack.Artist.Text
			LastFmPollingData.NowPlaying.LastFmURL = lastTrack.URL
			LastFmPollingData.NowPlaying.CoverArtURL = lastTrack.Image[len(lastTrack.Image)-1].Text

			// update when the last update took place
			LastFmPollingData.lastUpdated = time.Now()

			// update the pollCount
			LastFmPollingData.pollCount++
			LastFmPollCounter.Inc()

			LastFmPollingData.mu.Unlock()
		}

		time.Sleep(viper.GetDuration("lastfm.pollRateSecond") * time.Second)
	}
}
