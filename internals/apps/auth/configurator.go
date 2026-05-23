package auth

import (
	"github.com/halimdotnet/lulusiango/pkg/httpx"
	"github.com/labstack/echo/v5"
	"go.uber.org/fx"
)

// Controller defines the interface for HTTP route controllers.
type Controller interface {
	Register(e *echo.Echo)
}

// RouteParams groups the dependencies required to register routes using fx.In.
type RouteParams struct {
	fx.In

	Engine      httpx.Engine
	Controllers []Controller `group:"controllers"`
}

// RegisterRoutes registers all provided Controllers to the Echo instance.
func RegisterRoutes(params RouteParams) {
	e := params.Engine.Instance()
	for _, ctrl := range params.Controllers {
		ctrl.Register(e)
	}
}
