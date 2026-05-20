package rest

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/shared/config"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"go.uber.org/fx"
)

var Module = fx.Module("http_engine",
	fx.Provide(
		provideHttpConfig,
		NewEngine,
	),
	fx.Invoke(registerHooks),
)

func provideHttpConfig() (*HttpOption, error) {
	return config.BindJSONKey[*HttpOption]("server")
}

func registerHooks(lc fx.Lifecycle, e Engine, log logger.Logger, option *HttpOption) {
	var (
		serverCancel context.CancelFunc
		serverDone   = make(chan error, 1)
	)

	lc.Append(
		fx.Hook{
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
		},
	)
}
