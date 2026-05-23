package auth

import (
	db "github.com/halimdotnet/lulusiango/database/sqlc"
	loginV1 "github.com/halimdotnet/lulusiango/internals/apps/auth/features/login/v1"
	registerV1 "github.com/halimdotnet/lulusiango/internals/apps/auth/features/register/v1"
	"github.com/halimdotnet/lulusiango/internals/shared/fxapp"
	"github.com/halimdotnet/lulusiango/internals/shared/fxmodule"
	"github.com/halimdotnet/lulusiango/pkg/pgsql"
	"github.com/halimdotnet/lulusiango/pkg/validator"
	"go.uber.org/fx"
)

func NewApp() fxapp.Application {
	f := fxapp.NewBuilder()

	f.WithOptions(
		fxmodule.ModuleApp,
		fxmodule.ModuleLogger,
		fxmodule.ModulePgsql,
		fxmodule.ModuleServer,

		// Register V1 feature module
		registerV1.Module,

		// Login V1 feature module
		loginV1.Module,
	)

	f.WithProviders(
		// Provide application-wide validator
		validator.NewValidator,

		// Provide database queries wrapping the Pgsql client pool
		func(client pgsql.Client) *db.Queries {
			return db.New(client.Pool())
		},

		// Adapt concrete RegisterController to local Controller interface in the controllers group
		fx.Annotate(
			func(ctrl *registerV1.RegisterController) Controller {
				return ctrl
			},
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),

		// Adapt concrete LoginController to local Controller interface in the controllers group
		fx.Annotate(
			func(ctrl *loginV1.LoginController) Controller {
				return ctrl
			},
			fx.As(new(Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	)

	f.WithInvokes(
		// Register all controllers in the group
		RegisterRoutes,
	)

	return f.Build()
}
