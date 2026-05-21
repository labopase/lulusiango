package createuser

import (
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/models"
)

type CreateUserDtoRequest struct {
	Email     string            `json:"email" validate:"required,email"`
	Password  string            `json:"password" validate:"required,min=6"`
	Fullname  string            `json:"fullname" validate:"required"`
	Status    models.UserStatus `json:"status" validate:"required,oneof=active pending suspended"`
	CreatedBy *string           `json:"created_by"`
}

type CreateUserDTOResponse struct {
	ID string `json:"id"`
}
