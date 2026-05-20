package config

import (
	"errors"
	"log"
	"slices"

	"github.com/halimdotnet/lulusiango/internals/shared/constants"
	"go.uber.org/fx"
)

var Module = fx.Module("app_config",
	fx.Provide(
		func() (*AppConfig, error) {
			cfg, err := BindJSONKey[*AppConfig]("app")
			if err != nil {
				return nil, err
			}

			if cfg.Name == "" {
				return nil, errors.New("app.name is required")
			}

			if cfg.Version == "" {
				return nil, errors.New("app.version is required")
			}

			if !slices.Contains(constants.ListEnv, cfg.Environment) {
				return nil, errors.New("app.environment is invalid")
			}

			return cfg, nil
		},
	),
	fx.Invoke(
		func(cfg *AppConfig) {
			log.Printf("App Name: %s", cfg.Name)
			log.Printf("App Version: %s", cfg.Version)
			log.Printf("App Environment: %s", cfg.Environment)
			log.Printf("App Debug: %v", cfg.Debug)
		},
	),
)
