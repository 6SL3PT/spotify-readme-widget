package handlers

import (
	"bytes"

	"github.com/6sl3pt/spotify-readme-widget/services"
	"github.com/6sl3pt/spotify-readme-widget/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type SpotifyHandler struct {
	SpotifyService services.ISpotifyService
}

func NewSpotifyHandler(ss services.ISpotifyService) *SpotifyHandler {
	return &SpotifyHandler{
		SpotifyService: ss,
	}
}

func (h SpotifyHandler) HandlerShowWidget(c echo.Context) error  {
	track, err := h.SpotifyService.GetTrack()	
	if err != nil {
		cmp := views.Index(track, false, err.Error())
		return h.View(c, cmp)
	}

	cmp := views.Index(track, true, "")	
	return h.View(c, cmp)
}

func (h SpotifyHandler) View(c echo.Context, cmp templ.Component) error {
	var buf bytes.Buffer

	// Render to buffer
	err := cmp.Render(c.Request().Context(), &buf)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to render SVG: "+err.Error())
	}

	// Set proper content type for SVG blob
	c.Response().Header().Set(echo.HeaderContentType, "image/svg+xml")
	c.Response().WriteHeader(200)

	// Write the buffer to response
	_, err = buf.WriteTo(c.Response().Writer)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to write SVG to response: "+err.Error())
	}

	return nil
}
