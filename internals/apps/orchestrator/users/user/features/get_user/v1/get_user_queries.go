package getuser

import (
	"context"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/repository"
	"github.com/halimdotnet/lulusiango/internals/shared/utilities/validator"
)

type getUserQuery struct {
	userRepo  repository.PostgreUserRepository
	validator *validator.Validator
}

type GetUserQueries interface {
	GetUserByID(ctx context.Context, req *GetUserDtoRequest) (*GetUserDTOResponse, error)
	GetUserByEmail(ctx context.Context, req *GetUserDtoRequest) (*GetUserDTOResponse, error)
}

func NewGetUserQueries(userRepo repository.PostgreUserRepository) *getUserQuery {
	return &getUserQuery{userRepo: userRepo}
}

func (q *getUserQuery) GetUserByID(ctx context.Context, req *GetUserDtoRequest) (*GetUserDTOResponse, error) {
	return nil, nil
}

func (q *getUserQuery) GetUserByEmail(ctx context.Context, req *GetUserDtoRequest) (*GetUserDTOResponse, error) {
	return nil, nil
}

func (q *getUserQuery) validate(req *GetUserDtoRequest) error {

	return nil
}
