package v1

type CreateUserDtoRequest struct {
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=6"`
	Fullname  string  `json:"fullname" validate:"required"`
	Status    string  `json:"status" validate:"required,oneof=active pending suspended"`
	CreatedBy *string `json:"created_by"`
}

type CreateUserDTOResponse struct {
	ID string `json:"id"`
}
