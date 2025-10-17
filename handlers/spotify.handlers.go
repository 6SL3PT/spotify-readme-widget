package handlers

import (
	"reflect"

	"github.com/6sl3pt/spotify-readme-widget/services"
	"github.com/6sl3pt/spotify-readme-widget/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type SpotifyService interface {
	GetTrack() (services.Track, error)
}

type SpotifyHandler struct {
	SpotifyServices SpotifyService
}

func NewSpotifyHandler(ss SpotifyService) *SpotifyHandler {
	return &SpotifyHandler{
		SpotifyServices: ss,
	}
}

func (h SpotifyHandler) HandlerShowWidget(c echo.Context) error  {
	track, err := h.SpotifyServices.GetTrack()	
	if err != nil || reflect.DeepEqual(track, services.Track{}) {
		cmp := views.Index(track, false)
		return h.View(c, cmp)
	}

	cmp := views.Index(track, true)	
	return h.View(c, cmp)
}

func (h SpotifyHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
