package updateuser

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/halimdotnet/lulusiango/internals/shared/utilities/validator"
	"github.com/labstack/echo/v5"
)

type UpdateUserEndpoints struct {
	updateUser UpdateUser
	validator  *validator.Validator
}

func NewUpdateUserEndpoints(updateUser UpdateUser, validator *validator.Validator) *UpdateUserEndpoints {
	return &UpdateUserEndpoints{
		updateUser: updateUser,
		validator:  validator,
	}
}

func (e *UpdateUserEndpoints) Register(engine rest.Engine) {
	engine.Instance().PUT("api/v1/users/:id", e.UpdateUserHandler)
}

func (e *UpdateUserEndpoints) UpdateUserHandler(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "invalid id format",
		})
	}

	req := &UpdateUserDtoRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "failed to bind request",
		})
	}
	req.ID = id

	if errs := e.validator.ValidateStruct(req); len(errs) > 0 {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "validation error",
			Details: []interface{}{
				errs,
			},
		})
	}

	resp, err := e.updateUser.Execute(context.Background(), req)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, rest.Response{
				Error:   true,
				Message: "user not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, rest.Response{
			Error:   true,
			Message: "failed to update user",
			Details: []interface{}{
				err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, rest.Response{
		Error: false,
		Data:  resp,
	})
}
