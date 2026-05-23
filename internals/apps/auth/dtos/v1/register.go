package v1

type RegisterDtoRequest struct {
	Email           string `json:"email" validate:"required,email"`
	FullName        string `json:"fullname" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type RegisterDtoResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
