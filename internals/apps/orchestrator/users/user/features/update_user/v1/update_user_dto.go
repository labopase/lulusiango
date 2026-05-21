package updateuser

import (
	"github.com/google/uuid"
	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/models"
)

type UpdateUserDtoRequest struct {
	ID        uuid.UUID         `json:"id" validate:"required,uuid"`
	Email     string            `json:"email" validate:"omitempty,email"`
	Password  string            `json:"password" validate:"omitempty,min=6"`
	Fullname  string            `json:"fullname" validate:"omitempty"`
	Status    models.UserStatus `json:"status" validate:"omitempty,oneof=active pending suspended"`
	UpdatedBy *string           `json:"updated_by"`
}

type UpdateUserDTOResponse struct {
	ID string `json:"id"`
}
