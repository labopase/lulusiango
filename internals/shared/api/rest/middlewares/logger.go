package middlewares

import (
	"net/http"
	"time"

	"github.com/halimdotnet/lulusiango/internals/shared/constants"
	"github.com/halimdotnet/lulusiango/internals/shared/observability/logger"
	"github.com/labstack/echo/v5"
)

type LoggerConfig struct {
	Skipper     Skipper
	ServiceName string
}

func Logger(log logger.Logger, config LoggerConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultSkipper
	}

	now := time.Now

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			if config.Skipper(c) {
				return next(c)
			}

			err := next(c)

			if err != nil {
				panic(err)
			}

			start := now()
			req := c.Request()
			res := c.Response()

			resp, status := echo.ResolveResponseStatus(res, err)

			level := constants.LogInfo
			if status >= http.StatusInternalServerError {
				level = constants.LogError
			} else if status >= http.StatusBadRequest {
				level = constants.LogWarn
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			p := req.URL.Path
			if p == "" {
				p = "/"
			}

			field := []logger.Field{
				{
					Key:   "time",
					Value: start,
				},
				{
					Key:   "level",
					Value: level,
				},
				{
					Key:   "status",
					Value: status,
				},
				{
					Key:   "protocol",
					Value: req.Proto,
				},
				{
					Key:   "method",
					Value: req.Method,
				},
				{
					Key:   "uri",
					Value: req.RequestURI,
				},
				{
					Key:   "host",
					Value: req.Host,
				},
				{
					Key:   "uri_path",
					Value: p,
				},
				{
					Key:   "route_path",
					Value: c.Path(),
				},
				{
					Key:   "latency",
					Value: now().Sub(start),
				},
				{
					Key:   "remote_ip",
					Value: c.RealIP(),
				},
				{
					Key:   "request_id",
					Value: id,
				},
				{
					Key:   "user_agent",
					Value: req.UserAgent(),
				},
				{
					Key:   "bytes_out",
					Value: resp.Size,
				},
				{
					Key:   "byte_in",
					Value: req.ContentLength,
				},
			}

			if config.ServiceName != "" {
				field = append(field, logger.Field{
					Key:   "service",
					Value: config.ServiceName,
				})
			}

			log.Infow("REQUEST", field...)

			return nil
		}
	}
}
