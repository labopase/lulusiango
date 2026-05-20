package middlewares

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo/v5"
)

type RecoverConfig struct {
	Skipper           Skipper
	StackSize         int
	DisableStackAll   bool
	DisablePrintStack bool
}

type PanicStackError struct {
	Stack []byte
	Err   error
}

func (e *PanicStackError) Error() string {
	return fmt.Sprintf("[PANIC RECOVER] %s\n%s", e.Err.Error(), e.Stack)
}

func (e *PanicStackError) Unwrap() error {
	return e.Err
}

func Recover(config RecoverConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultSkipper
	}
	if config.StackSize <= 0 {
		config.StackSize = 4 << 10 // 4KB default
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {
					tmpErr, ok := r.(error)
					if !ok {
						tmpErr = fmt.Errorf("%v", r)
					}

					if !config.DisablePrintStack {
						stack := make([]byte, config.StackSize)
						length := runtime.Stack(stack, !config.DisableStackAll)
						tmpErr = &PanicStackError{
							Stack: stack[:length],
							Err:   tmpErr,
						}
					}
					err = tmpErr
				}
			}()

			return next(c)
		}
	}
}
