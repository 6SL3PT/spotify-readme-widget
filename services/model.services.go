package services

// Services
type SpotifyService struct {}

type ImageService struct {
	Blob []byte
}

// Data objects
type TrackSvg struct {
	Name      string
	Artist    string
	Image     string
}

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

// API Responses
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

