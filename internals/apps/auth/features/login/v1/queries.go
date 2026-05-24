package v1

import (
	"context"
	"errors"

	db "github.com/halimdotnet/lulusiango/database/sqlc"
	"github.com/halimdotnet/lulusiango/pkg/password"
)

type LoginUserQuery struct {
	Email    string
	Password string
}

type LoginUserQueryHandler interface {
	Handle(ctx context.Context, query LoginUserQuery) (db.Users, error)
}

type loginUserQueryHandler struct {
	queries *db.Queries
}

func NewLoginUserQueryHandler(queries *db.Queries) LoginUserQueryHandler {
	return &loginUserQueryHandler{
		queries: queries,
	}
}

func (h *loginUserQueryHandler) Handle(ctx context.Context, query LoginUserQuery) (db.Users, error) {
	user, err := h.queries.GetUserByEmail(ctx, query.Email)
	if err != nil {
		return db.Users{}, errors.New("invalid email or password")
	}

	err = password.VerifyBcrypt(query.Password, user.PasswordHash)
	if err != nil {
		return db.Users{}, errors.New("invalid email or password")
	}

	return user, nil
}
