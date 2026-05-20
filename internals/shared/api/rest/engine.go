package rest

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/halimdotnet/lulusiango/internals/shared/api/rest/middlewares"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"github.com/labstack/echo/v5"
)

type engine struct {
	echo   *echo.Echo
	log    logger.Logger
	option *HttpOption
}

func NewEngine(option *HttpOption, log logger.Logger) (Engine, error) {
	if option == nil {
		return nil, fmt.Errorf("httpx: option is required")
	}

	option.applyDefaults()

	e := echo.New()

	eng := &engine{
		echo:   e,
		log:    log,
		option: option,
	}

	eng.setupDefaultMiddleware()

	return eng, nil
}

func (e *engine) Start(ctx context.Context) error {

	ln, err := net.Listen("tcp", e.option.Addr())
	if err != nil {
		return fmt.Errorf("httpx: failed to listen on %s: %w", e.option.Addr(), err)
	}

	sc := echo.StartConfig{
		Address:         e.option.Addr(),
		Listener:        ln,
		HideBanner:      true,
		HidePort:        true,
		GracefulTimeout: e.option.ShutdownTimeout,
		BeforeServeFunc: func(s *http.Server) error {
			s.ReadTimeout = e.option.ReadTimeout
			s.IdleTimeout = e.option.IdleTimeout
			s.MaxHeaderBytes = e.option.MaxHeaderBytes
			return nil
		},
		OnShutdownError: func(err error) {
			e.log.Errorf("http server error on shutdown %s", err.Error())
		},
	}

	return sc.Start(ctx, e.echo)
}

func (e *engine) Instance() *echo.Echo {
	return e.echo
}

func (e *engine) setupDefaultMiddleware() {
	e.echo.Use(middlewares.Logger(e.log, middlewares.LoggerConfig{
		ServiceName: e.option.AppName,
	}))
	e.echo.Use(middlewares.Recover(middlewares.RecoverConfig{
		DisableStackAll:   false,
		DisablePrintStack: false,
		StackSize:         6 << 10, // 6KB
	}))
}
