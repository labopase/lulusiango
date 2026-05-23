package v1

import (
	"go.uber.org/fx"
)

// Module registers all components for the login feature in Uber Fx.
var Module = fx.Module("login_v1",
	fx.Provide(
		NewLoginUserQueryHandler,
		NewLoginController,
	),
)
