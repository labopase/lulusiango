package v1

import (
	"net/http"

	v1dto "github.com/halimdotnet/lulusiango/internals/apps/auth/dtos/v1"
	"github.com/halimdotnet/lulusiango/pkg/httpx"
	"github.com/halimdotnet/lulusiango/pkg/logger"
	"github.com/halimdotnet/lulusiango/pkg/validator"
	"github.com/labstack/echo/v5"
)

// RegisterController handles HTTP requests for user registration.
type RegisterController struct {
	handler RegisterUserCommandHandler
	logger  logger.Logger
	val     validator.Validator
}

// NewRegisterController creates a new RegisterController.
func NewRegisterController(
	handler RegisterUserCommandHandler,
	logger logger.Logger,
	val validator.Validator,
) *RegisterController {
	return &RegisterController{
		handler: handler,
		logger:  logger,
		val:     val,
	}
}

// Register registers the HTTP routes associated with this controller.
func (c *RegisterController) Register(e *echo.Echo) {
	e.POST("/api/v1/auth/register", c.RegisterUser)
}

// RegisterUser handles POST /api/v1/auth/register.
func (c *RegisterController) RegisterUser(ctx *echo.Context) error {
	var req v1dto.RegisterDtoRequest

	if err := ctx.Bind(&req); err != nil {
		c.logger.Errorw("failed to bind request", logger.Error(err))
		return ctx.JSON(http.StatusBadRequest, httpx.Response{
			Error:   true,
			Message: "invalid request body",
		})
	}

	// Validate the request payload
	if err := c.val.Validate(req); err != nil {
		c.logger.Errorw("validation failed", logger.Error(err))
		return ctx.JSON(http.StatusBadRequest, httpx.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	// Map DTO to CQRS Command
	cmd := RegisterUserCommand{
		Email:           req.Email,
		FullName:        req.FullName,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	}

	// Handle command execution
	_, err := c.handler.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		c.logger.Errorw("failed to register user", logger.Error(err))
		return ctx.JSON(http.StatusBadRequest, httpx.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	// Prepare response - tokens are empty as per constraint
	resp := v1dto.RegisterDtoResponse{
		Token:        "",
		RefreshToken: "",
	}

	return ctx.JSON(http.StatusCreated, httpx.Response{
		Error:   false,
		Message: "user registered successfully",
		Data:    resp,
	})
}
