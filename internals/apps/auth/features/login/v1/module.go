package v1

import (
	"go.uber.org/fx"
)

var Module = fx.Module("login_v1",
	fx.Provide(
		NewLoginUserQueryHandler,
		NewLoginController,
	),
)
