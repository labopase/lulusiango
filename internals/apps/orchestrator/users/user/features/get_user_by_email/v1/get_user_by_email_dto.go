package getuserbyemail

import (
	"time"

	"github.com/google/uuid"
)

type GetUserByEmailDtoRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetUserByEmailDTOResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *string   `json:"created_by"`
	UpdatedBy *string   `json:"updated_by"`
}
