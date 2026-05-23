package fxmodule

import (
	"github.com/halimdotnet/lulusiango/pkg/httpx"
	"github.com/halimdotnet/lulusiango/pkg/logger"
	"github.com/halimdotnet/lulusiango/pkg/pgsql"
	"github.com/halimdotnet/lulusiango/pkg/redis"
	"go.uber.org/fx"
)

var (
	ModuleApp = fx.Module("app_config",
		fx.Provide(provideAppOption),
	)

	ModuleLogger = fx.Module("logger",
		fx.Provide(
			provideLoggerOption,
			logger.NewLogger,
		),
		fx.Invoke(registerLoggerHooks),
	)

	ModuleRedis = fx.Module("redis",
		fx.Provide(
			provideRedisOption,
			redis.NewClient,
		),
		fx.Invoke(registerRedisHooks),
	)

	ModulePgsql = fx.Module("postgresql",
		fx.Provide(
			providePgsqlOption,
			pgsql.NewClient,
		),
		fx.Invoke(registerPgsqlHooks),
	)

	ModuleServer = fx.Module("http_server",
		fx.Provide(
			provideServerOption,
			httpx.NewEngine,
		),
		fx.Invoke(registerServerHooks),
	)
)
