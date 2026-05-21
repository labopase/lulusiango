package getuserbyemail

import (
	"context"
	"errors"
	"net/http"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/labstack/echo/v5"
)

type GetUserByEmailEndpoints struct {
	getUserByEmail GetUserByEmail
}

func NewGetUserByEmailEndpoints(getUserByEmail GetUserByEmail) *GetUserByEmailEndpoints {
	return &GetUserByEmailEndpoints{getUserByEmail: getUserByEmail}
}

func (e *GetUserByEmailEndpoints) Register(engine rest.Engine) {
	engine.Instance().GET("api/v1/users/email/:email", e.GetUserByEmailHandler)
}

func (e *GetUserByEmailEndpoints) GetUserByEmailHandler(c *echo.Context) error {
	emailParam := c.Param("email")
	if emailParam == "" {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "email is required",
		})
	}

	resp, err := e.getUserByEmail.Execute(context.Background(), &GetUserByEmailDtoRequest{
		Email: emailParam,
	})

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, rest.Response{
				Error:   true,
				Message: "user not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, rest.Response{
			Error:   true,
			Message: "internal server error",
			Details: []interface{}{
				err.Error(),
			},
		})
	}

	if resp == nil {
		return c.JSON(http.StatusNotFound, rest.Response{
			Error:   true,
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, rest.Response{
		Error: false,
		Data:  resp,
	})
}
