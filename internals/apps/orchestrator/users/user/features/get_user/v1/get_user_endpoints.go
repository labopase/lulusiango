package getuser

import (
	"time"

	"github.com/google/uuid"
)

type GetUserDtoRequest struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type GetUserDTOResponse struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedBy      uuid.UUID `json:"created_by"`
	UpdatedBy      uuid.UUID `json:"updated_by"`
}
