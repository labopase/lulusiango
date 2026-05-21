package createuser

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/models"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
)

type createUser struct {
	userRepo repository.UserRepositoryWriter
}

type CreateUser interface {
	Execute(ctx context.Context, req *CreateUserDtoRequest) (*CreateUserDTOResponse, error)
}

func NewCreateUser(userRepo repository.UserRepositoryWriter) CreateUser {
	return &createUser{userRepo: userRepo}
}

func (c *createUser) Execute(ctx context.Context, req *CreateUserDtoRequest) (*CreateUserDTOResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.Users{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Fullname:     req.Fullname,
		Status:       req.Status,
		CreatedBy:    req.CreatedBy,
		UpdatedBy:    req.CreatedBy,
	}

	result, err := c.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &CreateUserDTOResponse{
		ID: result.ID.String(),
	}, nil
}
