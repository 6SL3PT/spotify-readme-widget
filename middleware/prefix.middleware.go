package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// Middleware to removes the given prefix from the request URL path.
// This middleware is created to handle Vercel Functions deployment
// as the path registered in Echo is not match with serverless routes definition.
func StripPrefixMiddleware(prefix string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			// Strip prefix
			newPath := strings.TrimPrefix(req.URL.Path, prefix)
			if newPath == "" {
				newPath = "/"
			}

			req.URL.Path = newPath

			return next(c)
		}
	}
}
