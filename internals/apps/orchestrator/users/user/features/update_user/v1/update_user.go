package updateuser

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
)

type updateUser struct {
	userReader repository.UserRepositoryReader
	userWriter repository.UserRepositoryWriter
}

type UpdateUser interface {
	Execute(ctx context.Context, req *UpdateUserDtoRequest) (*UpdateUserDTOResponse, error)
}

func NewUpdateUser(userReader repository.UserRepositoryReader, userWriter repository.UserRepositoryWriter) UpdateUser {
	return &updateUser{
		userReader: userReader,
		userWriter: userWriter,
	}
}

func (c *updateUser) Execute(ctx context.Context, req *UpdateUserDtoRequest) (*UpdateUserDTOResponse, error) {
	user, err := c.userReader.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(hashedPassword)
	}
	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}
	if req.Status != "" {
		user.Status = req.Status
	}
	user.UpdatedBy = req.UpdatedBy

	result, err := c.userWriter.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &UpdateUserDTOResponse{
		ID: result.ID.String(),
	}, nil
}
