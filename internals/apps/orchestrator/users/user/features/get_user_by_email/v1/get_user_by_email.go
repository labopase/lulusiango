package getuserbyemail

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/models"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
)

type getUserByEmail struct {
	userRepo repository.UserRepositoryReader
}

type GetUserByEmail interface {
	Execute(ctx context.Context, req *GetUserByEmailDtoRequest) (*GetUserByEmailDTOResponse, error)
}

func NewGetUserByEmail(userRepo repository.UserRepositoryReader) GetUserByEmail {
	return &getUserByEmail{userRepo: userRepo}
}

func (q *getUserByEmail) Execute(ctx context.Context, req *GetUserByEmailDtoRequest) (*GetUserByEmailDTOResponse, error) {
	data, err := q.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return q.toDto(data), nil
}

func (q *getUserByEmail) toDto(user *models.Users) *GetUserByEmailDTOResponse {
	if user == nil {
		return nil
	}
	return &GetUserByEmailDTOResponse{
		ID:        user.ID,
		Email:     user.Email,
		Fullname:  user.Fullname,
		Status:    string(user.Status),
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
	}
}
