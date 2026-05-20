package healthz

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/halimdotnet/lulusiango/internals/shared/config"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"github.com/halimdotnet/lulusiango/internals/shared/store/pgsql"
	rds "github.com/halimdotnet/lulusiango/internals/shared/store/redis"
	"go.uber.org/fx"
)

var Module = fx.Module("health_checker",
	fx.Provide(
		func() (*HealthOptions, error) {
			return config.BindJSONKey[*HealthOptions]("health")
		},
		NewChecker,
		NewHandler,
	),
	fx.Invoke(
		func(
			lc fx.Lifecycle,
			ck Checker,
			rd rds.Client,
			h Handler,
			pg pgsql.Client,
			cfg *HealthOptions,
			e rest.Engine,
			log logger.Logger,
		) {
			log.Infow("Health Checker Initialized",
				logger.Bool("redis", cfg.Redis),
				logger.Bool("postgres", cfg.Postgres),
			)

			if cfg.Redis {
				ck.Register("redis", RedisCheck(rd))
			}
			if cfg.Postgres {
				ck.Register("postgres", PostgresCheck(pg))
			}

			route := e.Instance().Group("/health")
			route.GET("/live", h.LivenessHandler)
			route.GET("/ready", h.ReadinessHandler)
			route.GET("/startup", h.StartupHandler)

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					h.SetReady(true)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					h.SetReady(false)
					return nil
				},
			})
		},
	),
)
