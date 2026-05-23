package pgsql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type client struct {
	pool *pgxpool.Pool
}

func NewClient(option *PgsqlOption) (Client, error) {
	if option == nil {
		return nil, errors.New("pgsql: config is required")
	}

	option.applyDefaults()

	poolConfig, err := pgxpool.ParseConfig(option.Dsn())
	if err != nil {
		return nil, fmt.Errorf("pgsql: failed to parse config: %w", err)
	}

	poolConfig.MaxConns = option.MaxOpenConns
	poolConfig.MaxConnIdleTime = option.ConnMaxIdleTime
	poolConfig.MaxConnLifetime = option.ConnMaxLifetime
	poolConfig.HealthCheckPeriod = time.Minute
	poolConfig.ConnConfig.ConnectTimeout = option.ConnectTimeout

	p, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgsql: failed to create pool: %w", err)
	}

	if err := p.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("pgsql: failed to ping database: %w", err)
	}

	return &client{pool: p}, nil
}

func (c *client) Close() {
	c.pool.Close()
}

func (c *client) Pool() *pgxpool.Pool {
	return c.pool
}

func (c *client) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}
