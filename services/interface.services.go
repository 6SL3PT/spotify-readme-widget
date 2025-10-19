package services

type ISpotifyService interface {
	GetTrack() (TrackSvg, error)
}
