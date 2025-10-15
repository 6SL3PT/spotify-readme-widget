package main 

import (
	"github.com/6sl3pt/spotify-readme-widget/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	h := handlers.New()
	handlers.SetupRoutes(e, h)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
