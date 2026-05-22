package createuser

import (
	"context"
	"net/http"

	dto "github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/dtos/v1"
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/halimdotnet/lulusiango/internals/shared/utilities/validator"
	"github.com/labstack/echo/v5"
)

type CreateUserEndpoints struct {
	createUser CreateUser
	validator  *validator.Validator
}

func NewCreateUserEndpoints(createUser CreateUser, validator *validator.Validator) *CreateUserEndpoints {
	return &CreateUserEndpoints{
		createUser: createUser,
		validator:  validator,
	}
}

func (e *CreateUserEndpoints) Register(engine rest.Engine) {
	engine.Instance().POST("api/v1/users", e.CreateUserHandler)
}

func (e *CreateUserEndpoints) CreateUserHandler(c *echo.Context) error {
	req := &dto.CreateUserDtoRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "failed to bind request",
		})
	}

	if errs := e.validator.ValidateStruct(req); len(errs) > 0 {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "validation error",
			Details: []interface{}{
				errs,
			},
		})
	}

	resp, err := e.createUser.Execute(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rest.Response{
			Error:   true,
			Message: "failed to create user",
			Details: []interface{}{
				err.Error(),
			},
		})
	}

	return c.JSON(http.StatusCreated, rest.Response{
		Error: false,
		Data:  resp,
	})
}
