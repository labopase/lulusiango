package v1

import (
	"context"
	"errors"

	db "github.com/halimdotnet/lulusiango/database/sqlc"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUserCommand holds the payload for registering a new user.
type RegisterUserCommand struct {
	Email           string
	FullName        string
	Password        string
	ConfirmPassword string
}

// RegisterUserCommandHandler defines the interface for handling the registration command.
type RegisterUserCommandHandler interface {
	Handle(ctx context.Context, cmd RegisterUserCommand) (db.Users, error)
}

type registerUserCommandHandler struct {
	queries *db.Queries
}

// NewRegisterUserCommandHandler creates a new instance of RegisterUserCommandHandler.
func NewRegisterUserCommandHandler(queries *db.Queries) RegisterUserCommandHandler {
	return &registerUserCommandHandler{
		queries: queries,
	}
}

// Handle executes the registration command, performing validation, hashing, and database insertion.
func (h *registerUserCommandHandler) Handle(ctx context.Context, cmd RegisterUserCommand) (db.Users, error) {
	if cmd.Password != cmd.ConfirmPassword {
		return db.Users{}, errors.New("passwords do not match")
	}

	// Check if user already exists
	_, err := h.queries.GetUserByEmail(ctx, cmd.Email)
	if err == nil {
		return db.Users{}, errors.New("user with this email already exists")
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.Users{}, err
	}

	// Create user in the database
	user, err := h.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        cmd.Email,
		PasswordHash: string(hashedPassword),
		Fullname:     cmd.FullName,
		Status:       db.UserStatusPending,
		CreatedBy:    nil,
		UpdatedBy:    nil,
	})
	if err != nil {
		return db.Users{}, err
	}

	return user, nil
}
