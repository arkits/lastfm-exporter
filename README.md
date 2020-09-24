# LastFM Exporter

Polls LastFM for your currently listening tracks and exposes the data from a REST endpoint.

```
~/Dev/musick master ❯ https archit.xyz/musick/now
HTTP/1.1 200 OK
Connection: keep-alive
Content-Length: 322
Content-Type: application/json

{
    "nowPlaying": {
        "trackName": "PDA",
        "artistName": "Interpol",
        "albumName": "Turn On the Bright Lights: The Tenth Anniversary Edition (Remastered)",
        "coverArtUrl": "https://lastfm.freetls.fastly.net/i/u/300x300/6d9f06200d85475cbd970f3437997979.jpg",
        "lastFmUrl": "https://www.last.fm/music/Interpol/_/PDA"
    },
    "username": "architkhode"
}
```

## Development

### Config Management

Create a `config.yml` based on the `config.sample.yml` with your LastFM API key, secret, and target username.

### Running Locally

```bash
go mod download
go run main.go
```
