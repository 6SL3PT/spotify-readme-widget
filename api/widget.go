package api

import (
	"net/http"

	"github.com/6sl3pt/spotify-readme-widget/handlers"
	"github.com/6sl3pt/spotify-readme-widget/middleware"
	"github.com/6sl3pt/spotify-readme-widget/services"

	"github.com/labstack/echo/v4"
)

const vercelPrefix = "/api/widget"

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	// Apply middleware
	e.Pre(middleware.StripPrefixMiddleware(vercelPrefix))

	// Initialize
	ss := services.NewSpotifyService()
	h := handlers.NewSpotifyHandler(ss)
	
	// Setup route
	e.GET("/", h.HandlerShowWidget)

	// Start server
	e.ServeHTTP(w, r)
}
