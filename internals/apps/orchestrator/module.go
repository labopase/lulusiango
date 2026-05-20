package orchestrator

import (
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/halimdotnet/lulusiango/internals/shared/config"
	"github.com/halimdotnet/lulusiango/internals/shared/fxapp"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/healthz"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"github.com/halimdotnet/lulusiango/internals/shared/store/pgsql"
	rds "github.com/halimdotnet/lulusiango/internals/shared/store/redis"
)

func NewApp() fxapp.Application {
	return fxapp.NewBuilder().
		WithOptions(
			config.Module,
			logger.Module,
			pgsql.Module,
			rds.Module,
			healthz.Module,
			rest.Module,
		).
		Build()
}
