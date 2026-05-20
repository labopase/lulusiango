package logger

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/shared/config"
	"go.uber.org/fx"
)

var Module = fx.Module("logger",
	fx.Provide(
		provideLoggerConfig,
		NewLogger,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, log Logger, option *LoggerOption) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Infow("Logger initialized",
						String("environment", option.Environment),
						Bool("enable_caller", option.EnableCaller),
						Bool("enable_trace", option.EnableTrace),
					)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Info("Logger shutting down")
					_ = log.Sync()
					return nil
				},
			})
		},
	),
)

func provideLoggerConfig(appCfg *config.AppConfig) (*LoggerOption, error) {
	option, err := config.BindJSONKey[*LoggerOption]("logger")
	if err != nil {
		return nil, err
	}

	option.Environment = appCfg.Environment

	return option, nil
}
