package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	currentlyPlayingURL = "https://api.spotify.com/v1/me/player/currently-playing"
	recentlyPlayedURL   = "https://api.spotify.com/v1/me/player/recently-played?limit=1"
	tokenURL            = "https://accounts.spotify.com/api/token"
)

func NewSpotifyService() *SpotifyService {
	return &SpotifyService{}
}

func (ss SpotifyService) GetTrack() (TrackSvg, error) {
	accessToken, err := refreshToken()
	if err != nil {
		return TrackSvg{}, fmt.Errorf("Failed to refresh token: %w", err)
	}

	// Get currently play track
	if track, err := getCurrentlyPlayingTrack(accessToken); err == nil {
		return track, nil
	}

	// Get recently play track, in case there's no track playing at fetching time
	if track, err := getRecentlyPlayedTrack(accessToken); err == nil {
		return track, nil
	}

	return TrackSvg{}, errors.New("Failed to retrieve both currently and recently played tracks")
}

func getCurrentlyPlayingTrack(token string) (TrackSvg, error) {
	// Fetch data
	req, err := http.NewRequest("GET", currentlyPlayingURL, nil)
	if err != nil {
		return TrackSvg{}, fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer " + token)
	
	body, err := FetchApi(req)
	if err != nil {
		return TrackSvg{}, err
	}

	// Extract data
	var response CurrentlyPlayingResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return TrackSvg{}, fmt.Errorf("Failed to parse currently playing response: %w", err)
	}

	return mapTrackToTrackSvg(response.Item), nil
}

func getRecentlyPlayedTrack(token string) (TrackSvg, error) {
	// Fetch data
	req, err := http.NewRequest("GET", recentlyPlayedURL, nil)
	if err != nil {
		return TrackSvg{}, fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer " + token)
	
	body, err := FetchApi(req)
	if err != nil {
		return TrackSvg{}, err
	}

	// Extract data
	var response RecentlyPlayedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return TrackSvg{}, fmt.Errorf("Failed to parse recently played response: %w", err)
	}

	if len(response.Items) == 0 {
		return TrackSvg{}, errors.New("No recently played tracks found")
	}

	return mapTrackToTrackSvg(response.Items[0].Track), nil
}

func mapTrackToTrackSvg(track Track) TrackSvg {
	imageService := NewImageService(track.Album.Images[0].Url)
	b64Image, _ := imageService.GetBase64()

	return TrackSvg{Name: track.Name, Artist: track.Artists[0].Name, Image: b64Image}
}

func refreshToken() (string, error) {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	// Setup form data
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	authString := fmt.Sprintf("%s:%s", clientId, clientSecret)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	// Send request
	req, _ := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic " + encodedAuth)

	body, err := FetchApi(req)
	if err != nil {
		return "", err
	}

	// Extract data
	var response AccessTokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("Failed to parse access token response: %w", err)
	}

	return response.AccessToken, nil
}
