package deleteuser

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
	"github.com/halimdotnet/lulusiango/internals/shared/api/rest"
	"github.com/labstack/echo/v5"
)

type DeleteUserEndpoints struct {
	deleteUser DeleteUser
}

func NewDeleteUserEndpoints(deleteUser DeleteUser) *DeleteUserEndpoints {
	return &DeleteUserEndpoints{deleteUser: deleteUser}
}

func (e *DeleteUserEndpoints) Register(engine rest.Engine) {
	engine.Instance().DELETE("api/v1/users/:id", e.DeleteUserHandler)
}

func (e *DeleteUserEndpoints) DeleteUserHandler(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rest.Response{
			Error:   true,
			Message: "invalid id format",
		})
	}

	err = e.deleteUser.Execute(context.Background(), &DeleteUserDtoRequest{
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
			Message: "failed to delete user",
			Details: []interface{}{
				err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, rest.Response{
		Error:   false,
		Message: "user deleted successfully",
	})
}
