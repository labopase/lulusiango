package fxmodule

import (
	"fmt"
	"slices"

	appCfg "github.com/halimdotnet/lulusiango/internals/shared/config"
	"github.com/halimdotnet/lulusiango/internals/shared/constants"
	"github.com/halimdotnet/lulusiango/pkg/httpx"
	"github.com/halimdotnet/lulusiango/pkg/logger"
	"github.com/halimdotnet/lulusiango/pkg/pgsql"
	"github.com/halimdotnet/lulusiango/pkg/redis"
)

func provideAppOption() (*appCfg.AppConfig, error) {
	cfg, err := appCfg.BindJSONConfigKey[*appCfg.AppConfig]("app")
	if err != nil {
		return nil, fmt.Errorf("fx module: could not load app config: %w", err)
	}

	if !slices.Contains(constants.ListEnv, cfg.Environment) {
		return nil, fmt.Errorf("fx module: invalid environment '%s'", cfg.Environment)
	}

	return cfg, nil
}

func providePgsqlOption() (*pgsql.PgsqlOption, error) {
	cfg, err := appCfg.BindJSONConfigKey[*pgsql.PgsqlOption]("postgres")
	if err != nil {
		return nil, fmt.Errorf("fx module: could not load pgsql config: %w", err)
	}

	return cfg, nil
}

func provideRedisOption() (*redis.RedisOption, error) {
	cfg, err := appCfg.BindJSONConfigKey[*redis.RedisOption]("redis")
	if err != nil {
		return nil, fmt.Errorf("fx module: could not load redis config: %w", err)
	}

	return cfg, nil
}

func provideLoggerOption(app *appCfg.AppConfig) (*logger.LoggerOption, error) {
	cfg, err := appCfg.BindJSONConfigKey[*logger.LoggerOption]("logger")
	if err != nil {
		return nil, fmt.Errorf("fx module: could not load logger config: %w", err)
	}

	cfg.IsDevelopment = app.Environment == constants.EnvProduction || app.Environment == constants.EnvTest

	return cfg, nil
}

func provideServerOption() (*httpx.HttpOption, error) {
	cfg, err := appCfg.BindJSONConfigKey[*httpx.HttpOption]("server")
	if err != nil {
		return nil, fmt.Errorf("fx module: could not load server config: %w", err)
	}

	return cfg, nil
}
