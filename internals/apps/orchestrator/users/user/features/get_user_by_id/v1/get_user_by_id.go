package getuserbyid

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/models"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
)

type getUserByID struct {
	userRepo repository.UserRepositoryReader
}

type GetUserByID interface {
	Execute(ctx context.Context, req *GetUserByIDDtoRequest) (*GetUserByIDDTOResponse, error)
}

func NewGetUserByID(userRepo repository.UserRepositoryReader) GetUserByID {
	return &getUserByID{userRepo: userRepo}
}

func (q *getUserByID) Execute(ctx context.Context, req *GetUserByIDDtoRequest) (*GetUserByIDDTOResponse, error) {
	data, err := q.userRepo.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return q.toDto(data), nil
}

func (q *getUserByID) toDto(user *models.Users) *GetUserByIDDTOResponse {
	if user == nil {
		return nil
	}
	return &GetUserByIDDTOResponse{
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
