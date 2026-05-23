package v1

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	db "github.com/halimdotnet/lulusiango/database/sqlc"
)

// LoginUserQuery holds the payload to execute user credential verification.
type LoginUserQuery struct {
	Email    string
	Password string
}

// LoginUserQueryHandler defines the interface for handling the login query.
type LoginUserQueryHandler interface {
	Handle(ctx context.Context, query LoginUserQuery) (db.Users, error)
}

type loginUserQueryHandler struct {
	queries *db.Queries
}

// NewLoginUserQueryHandler creates a new instance of LoginUserQueryHandler.
func NewLoginUserQueryHandler(queries *db.Queries) LoginUserQueryHandler {
	return &loginUserQueryHandler{
		queries: queries,
	}
}

// Handle executes the login query by fetching the user and comparing the hashed password.
func (h *loginUserQueryHandler) Handle(ctx context.Context, query LoginUserQuery) (db.Users, error) {
	// Fetch user by email
	user, err := h.queries.GetUserByEmail(ctx, query.Email)
	if err != nil {
		// Non-disclosing generic error for security
		return db.Users{}, errors.New("invalid email or password")
	}

	// Compare stored password hash with entered password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(query.Password))
	if err != nil {
		// Non-disclosing generic error for security
		return db.Users{}, errors.New("invalid email or password")
	}

	return user, nil
}
