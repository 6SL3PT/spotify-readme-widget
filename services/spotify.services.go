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

type Track struct {
	Name      string `json:"name"`
	Artists []struct {
		Name    string `json:"name"`
	} `json:"artists"`
	Album     struct {
		Images []struct {
			Url     string `json:"url"`
		} `json:"images"`
	} `json:"album"`
}

type CurrentlyPlayingResponse struct {
	Item      Track  `json:"item"`
}

type RecentlyPlayedResponse struct {
	Items []struct {
		Track   Track  `json:"track"`
	} `json:"items"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ApiService interface {
	FetchApi(req *http.Request) ([]byte, error)
}

type SpotifyServices struct {
	ApiServices ApiService
}

func NewSpotifyService(as ApiService) *SpotifyServices {
	return &SpotifyServices{
		ApiServices: as,
	}
}

func (ss SpotifyServices) GetTrack() (Track, error) {
	accessToken, err := ss.refreshToken()
	if err != nil {
		return Track{}, fmt.Errorf("Failed to refresh token: %w", err)
	}

	// Get currently play track
	if track, err := ss.getCurrentlyPlayingTrack(accessToken); err == nil {
		return track, nil
	}

	// Get recently play track, in case there's no track playing at fetching time
	if track, err := ss.getRecentlyPlayedTrack(accessToken); err == nil {
		return track, nil
	}

	return Track{}, errors.New("Failed to retrieve both currently and recently played tracks")
}

func (ss SpotifyServices) getCurrentlyPlayingTrack(token string) (Track, error) {
	// Fetch data
	req, err := http.NewRequest("GET", currentlyPlayingURL, nil)
	if err != nil {
		return Track{}, fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer " + token)
	
	body, err := ss.ApiServices.FetchApi(req)
	if err != nil {
		return Track{}, err
	}

	// Extract data
	var response CurrentlyPlayingResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return Track{}, fmt.Errorf("Failed to parse currently playing response: %w", err)
	}

	return response.Item, nil
}

func (ss SpotifyServices) getRecentlyPlayedTrack(token string) (Track, error) {
	// Fetch data
	req, err := http.NewRequest("GET", recentlyPlayedURL, nil)
	if err != nil {
		return Track{}, fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer " + token)
	
	body, err := ss.ApiServices.FetchApi(req)
	if err != nil {
		return Track{}, err
	}

	// Extract data
	var response RecentlyPlayedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return Track{}, fmt.Errorf("Failed to parse recently played response: %w", err)
	}

	if len(response.Items) == 0 {
		return Track{}, errors.New("No recently played tracks found")
	}

	return response.Items[0].Track, nil
}

func (ss SpotifyServices) refreshToken() (string, error) {
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

	body, err := ss.ApiServices.FetchApi(req)
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
