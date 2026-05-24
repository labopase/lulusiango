package v1

import (
	"net/http"

	v1dto "github.com/halimdotnet/lulusiango/internals/apps/auth/dtos/v1"
	"github.com/halimdotnet/lulusiango/pkg/httpx"
	"github.com/halimdotnet/lulusiango/pkg/logger"
	"github.com/halimdotnet/lulusiango/pkg/validator"
	"github.com/labstack/echo/v5"
)

type LoginController struct {
	handler LoginUserQueryHandler
	logger  logger.Logger
	val     validator.Validator
}

func NewLoginController(
	handler LoginUserQueryHandler,
	logger logger.Logger,
	val validator.Validator,
) *LoginController {
	return &LoginController{
		handler: handler,
		logger:  logger,
		val:     val,
	}
}

func (c *LoginController) Register(e *echo.Echo) {
	e.POST("/api/v1/auth/login", c.LoginUser)
}

func (c *LoginController) LoginUser(ctx *echo.Context) error {
	var req v1dto.LoginDtoRequest

	if err := ctx.Bind(&req); err != nil {
		c.logger.Errorw("failed to bind request", logger.Error(err))
		return ctx.JSON(http.StatusBadRequest, httpx.Response{
			Error:   true,
			Message: "invalid request body",
		})
	}

	if err := c.val.Validate(req); err != nil {
		c.logger.Errorw("validation failed", logger.Error(err))
		return ctx.JSON(http.StatusBadRequest, httpx.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	query := LoginUserQuery{
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := c.handler.Handle(ctx.Request().Context(), query)
	if err != nil {
		c.logger.Errorw("failed to authenticate user", logger.Error(err))
		return ctx.JSON(http.StatusUnauthorized, httpx.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	resp := v1dto.LoginDtoResponse{
		Token:        "",
		RefreshToken: "",
	}

	return ctx.JSON(http.StatusOK, httpx.Response{
		Error:   false,
		Message: "login successful",
		Data:    resp,
	})
}
