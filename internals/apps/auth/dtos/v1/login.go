package v1

type LoginDtoRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginDtoResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
