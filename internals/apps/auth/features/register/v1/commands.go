package v1

import (
	"context"
	"errors"

	db "github.com/halimdotnet/lulusiango/database/sqlc"
	"github.com/halimdotnet/lulusiango/pkg/password"
)

type RegisterUserCommand struct {
	Email           string
	FullName        string
	Password        string
	ConfirmPassword string
}

type RegisterUserCommandHandler interface {
	Handle(ctx context.Context, cmd RegisterUserCommand) (db.Users, error)
}

type registerUserCommandHandler struct {
	queries *db.Queries
}

func NewRegisterUserCommandHandler(queries *db.Queries) RegisterUserCommandHandler {
	return &registerUserCommandHandler{
		queries: queries,
	}
}

func (h *registerUserCommandHandler) Handle(ctx context.Context, cmd RegisterUserCommand) (db.Users, error) {
	if cmd.Password != cmd.ConfirmPassword {
		return db.Users{}, errors.New("passwords do not match")
	}

	_, err := h.queries.GetUserByEmail(ctx, cmd.Email)
	if err == nil {
		return db.Users{}, errors.New("user with this email already exists")
	}

	hashedPassword, err := password.HashBcrypt(cmd.Password)
	if err != nil {
		return db.Users{}, err
	}

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
