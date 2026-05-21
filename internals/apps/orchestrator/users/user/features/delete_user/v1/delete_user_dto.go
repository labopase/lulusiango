package deleteuser

import (
	"github.com/google/uuid"
)

type DeleteUserDtoRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
