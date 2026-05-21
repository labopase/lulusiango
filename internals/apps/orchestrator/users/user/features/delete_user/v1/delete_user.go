package deleteuser

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
)

type deleteUser struct {
	userReader repository.UserRepositoryReader
	userWriter repository.UserRepositoryWriter
}

type DeleteUser interface {
	Execute(ctx context.Context, req *DeleteUserDtoRequest) error
}

func NewDeleteUser(userReader repository.UserRepositoryReader, userWriter repository.UserRepositoryWriter) DeleteUser {
	return &deleteUser{
		userReader: userReader,
		userWriter: userWriter,
	}
}

func (c *deleteUser) Execute(ctx context.Context, req *DeleteUserDtoRequest) error {
	_, err := c.userReader.GetUserByID(ctx, req.ID)
	if err != nil {
		return err
	}

	return c.userWriter.DeleteUser(ctx, req.ID)
}
