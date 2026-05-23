package v1

import (
	"go.uber.org/fx"
)

// Module registers all components for the register feature in Uber Fx.
var Module = fx.Module("register_v1",
	fx.Provide(
		NewRegisterUserCommandHandler,
		NewRegisterController,
	),
)
