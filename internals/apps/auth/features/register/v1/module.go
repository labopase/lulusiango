package v1

import (
	"go.uber.org/fx"
)

var Module = fx.Module("register_v1",
	fx.Provide(
		NewRegisterUserCommandHandler,
		NewRegisterController,
	),
)
