package pgsql

import (
	"fmt"
	"time"
)

type PgsqlOption struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"dbname"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	ConnectTimeout  time.Duration `mapstructure:"connect_timeout"`
	ApplicationName string        `mapstructure:"application_name"`

	MaxOpenConns    int32         `mapstructure:"max_open_conns"`
	MaxIdleConns    int32         `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

func (o *PgsqlOption) Dsn() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s&connect_timeout=%.0f&application_name=%s", o.User, o.Password, o.Host, o.Port, o.DBName, o.SSLMode, o.ConnectTimeout.Seconds(), o.ApplicationName)
}

func (o *PgsqlOption) applyDefaults() {
	if o.ConnectTimeout == 0 {
		o.ConnectTimeout = 30 * time.Second
	}
	if o.MaxOpenConns == 0 {
		o.MaxOpenConns = 10
	}
	if o.MaxIdleConns == 0 {
		o.MaxIdleConns = 5
	}
	if o.ConnMaxLifetime == 0 {
		o.ConnMaxLifetime = 1 * time.Hour
	}
	if o.ConnMaxIdleTime == 0 {
		o.ConnMaxIdleTime = 1 * time.Hour
	}
}
