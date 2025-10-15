package handlers

import (
	"github.com/6sl3pt/spotify-readme-widget/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type SpotifyHandler struct {
	// Services
}

func New() *SpotifyHandler {
	return &SpotifyHandler{
		// Services
	}
}

func (h *SpotifyHandler) HandlerShowWidget(c echo.Context) error  {
	cmp := views.Index("Test Title", "Hello, this is test message")
	
	return h.View(c, cmp)
}

func (h *SpotifyHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
