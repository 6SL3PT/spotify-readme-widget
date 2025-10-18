package api

import (
	"net/http"
	"strings"

	"github.com/6sl3pt/spotify-readme-widget/handlers"
	"github.com/6sl3pt/spotify-readme-widget/services"

	"github.com/labstack/echo/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	// Initialize
	s := services.NewSpotifyService()
	h := handlers.NewSpotifyHandler(s)
	
	// Setup route
	e.GET("/", h.HandlerShowWidget)

	// Strip Vercel serverless prefix
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/widget")
	if r.URL.Path == "" {
		r.URL.Path = "/"
	}


	// Start server
	e.ServeHTTP(w, r)
}
