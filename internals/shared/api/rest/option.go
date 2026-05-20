package rest

import (
	"fmt"
	"time"
)

const (
	writeTimeout    = 15 * time.Second
	readTimeout     = 15 * time.Second
	idleTimeout     = 30 * time.Second
	shutdownTimeout = 30 * time.Second
	maxHeaderBytes  = 1 << 20 // 1MB
	maxBodyBytes    = 1 << 20
)

type HttpOption struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`

	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	MaxHeaderBytes  int           `mapstructure:"max_header_bytes"`
	MaxBodyBytes    int           `mapstructure:"max_body_bytes"`
}

func (o *HttpOption) Addr() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o *HttpOption) applyDefaults() {
	if o.WriteTimeout == 0 {
		o.WriteTimeout = writeTimeout
	}
	if o.ReadTimeout == 0 {
		o.ReadTimeout = readTimeout
	}
	if o.IdleTimeout == 0 {
		o.IdleTimeout = idleTimeout
	}
	if o.ShutdownTimeout == 0 {
		o.ShutdownTimeout = shutdownTimeout
	}
	if o.MaxHeaderBytes == 0 {
		o.MaxHeaderBytes = maxHeaderBytes
	}
	if o.MaxBodyBytes == 0 {
		o.MaxBodyBytes = maxBodyBytes
	}
}
