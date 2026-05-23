package auth

import (
	"github.com/halimdotnet/lulusiango/internals/shared/fxapp"
	"github.com/halimdotnet/lulusiango/internals/shared/fxmodule"
)

func NewApp() fxapp.Application {
	f := fxapp.NewBuilder()

	f.WithOptions(
		fxmodule.ModuleApp,
		fxmodule.ModuleLogger,
		fxmodule.ModulePgsql,
		fxmodule.ModuleServer,
	)

	return f.Build()
}
