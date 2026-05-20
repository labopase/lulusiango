package middlewares

import (
	"strings"

	"github.com/labstack/echo/v5"
)

type Skipper func(c *echo.Context) bool

func DefaultSkipper(e *echo.Context) bool {
	return strings.Contains(e.Request().URL.Path, "swagger") ||
		strings.Contains(e.Request().URL.Path, "metrics") ||
		strings.Contains(e.Request().URL.Path, "health") ||
		strings.Contains(e.Request().URL.Path, "debug") ||
		strings.Contains(e.Request().URL.Path, "pprof")
}
