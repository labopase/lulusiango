package redis

import (
	"context"
	"strconv"

	"github.com/halimdotnet/lulusiango/internals/shared/config"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module("redis",
	fx.Provide(
		func() (*RedisOption, error) {
			return config.BindJSONKey[*RedisOption]("redis")
		},
		NewClient,
		func(c Client) redis.UniversalClient {
			return c.Rds()
		},
	),
	fx.Invoke(
		func(lc fx.Lifecycle, c Client, log logger.Logger, option *RedisOption) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Infow("Redis connected",
							logger.Any("addrs", option.Addrs),
							logger.String("pool_size", strconv.Itoa(option.PoolSize)),
							logger.String("min_idle_conns", strconv.Itoa(option.MinIdleConns)),
							logger.String("dial_timeout", option.DialTimeout.String()),
							logger.String("read_timeout", option.ReadTimeout.String()),
							logger.String("write_timeout", option.WriteTimeout.String()),
						)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						log.Info("Redis disconnected")
						return c.Close()
					},
				},
			)
		},
	),
)
