package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	poolSize     = 10
	minIdleConns = 5
	dialTimeout  = 5 * time.Second
	readTimeout  = 3 * time.Second
	writeTimeout = 3 * time.Second
)

type RedisOption struct {
	Addrs        []string      `mapstructure:"addrs"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	PoolSize     int           `mapstructure:"pool_size"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func (o *RedisOption) redisOption() *redis.UniversalOptions {
	opt := &redis.UniversalOptions{
		Addrs:        o.Addrs,
		Password:     o.Password,
		PoolSize:     o.PoolSize,
		MinIdleConns: o.MinIdleConns,
		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	}

	if len(o.Addrs) == 1 {
		opt.DB = o.DB
	}

	return opt
}

func (o *RedisOption) applyDefaults() {
	if len(o.Addrs) == 0 {
		o.Addrs = []string{"127.0.0.1:6379"}
	}

	if o.PoolSize == 0 {
		o.PoolSize = poolSize
	}

	if o.MinIdleConns == 0 {
		o.MinIdleConns = minIdleConns
	}

	if o.DialTimeout == 0 {
		o.DialTimeout = dialTimeout
	}

	if o.ReadTimeout == 0 {
		o.ReadTimeout = readTimeout
	}

	if o.WriteTimeout == 0 {
		o.WriteTimeout = writeTimeout
	}
}
