package pgsql

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/shared/config"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"go.uber.org/fx"
)

var Module = fx.Module("postgres",
	fx.Provide(
		func() (*PgsqlOption, error) {
			return config.BindJSONKey[*PgsqlOption]("postgres")
		},
		NewClient,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, c Client, log logger.Logger, option *PgsqlOption) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Infow("PostgreSQL connected", logger.String("host", option.Host), logger.Int("port", option.Port))
						return nil
					},
					OnStop: func(ctx context.Context) error {
						c.Close()
						log.Info("PostgreSQL disconnected")
						return nil
					},
				},
			)
		},
	),
)
