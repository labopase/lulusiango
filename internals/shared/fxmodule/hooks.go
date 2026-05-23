package fxmodule

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/shared/httpx"
	"github.com/halimdotnet/lulusiango/pkg/logger"
	"github.com/halimdotnet/lulusiango/pkg/pgsql"
	"github.com/halimdotnet/lulusiango/pkg/redis"
	"go.uber.org/fx"
)

func registerPgsqlHooks(lc fx.Lifecycle, pg pgsql.Client) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pg.Close()
			return nil
		},
	})
}

func registerRedisHooks(lc fx.Lifecycle, rds redis.Client) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return rds.Close()
		},
	})
}

func registerLoggerHooks(lc fx.Lifecycle, log logger.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = log.Sync()
			return nil
		},
	})
}

func registerServerHooks(lc fx.Lifecycle, e httpx.Engine, log logger.Logger, option *httpx.HttpOption) {
	var (
		serverCancel context.CancelFunc
		serverDone   = make(chan error, 1)
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Infow("initializing http server",
				logger.String("addr", option.Addr()),
				logger.String("read_timeout", option.ReadTimeout.String()),
				logger.String("idle_timeout", option.IdleTimeout.String()),
				logger.String("shutdown_timeout", option.ShutdownTimeout.String()),
			)
			serverCtx, cancel := context.WithCancel(context.Background())
			serverCancel = cancel

			go func() {
				serverDone <- e.Start(serverCtx)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Infow("shutting down http server",
				logger.String("addr", option.Addr()),
			)

			serverCancel()

			select {
			case err := <-serverDone:
				return err
			case <-ctx.Done():
				return ctx.Err()
			}
		},
	})
}
