package getuserbyid

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/labstack/echo/v5"
)

type GetUserByIDEndpoints struct {
	getUserByID GetUserByID
}

func NewGetUserByIDEndpoints(getUserByID GetUserByID) *GetUserByIDEndpoints {
	return &GetUserByIDEndpoints{getUserByID: getUserByID}
}

func (e *GetUserByIDEndpoints) Register(engine rest.Engine) {
	engine.Instance().GET("api/v1/users/:id", e.GetUserByIDHandler)
}

func (e *GetUserByIDEndpoints) GetUserByIDHandler(c *echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "invalid id format",
		})
	}

	resp, err := e.getUserByID.Execute(context.Background(), &GetUserByIDDtoRequest{
		ID: id,
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
