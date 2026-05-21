package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/halimdotnet/lulusiango/internals/apps/orchestrator/users/user/models"
	"github.com/halimdotnet/lulusiango/internals/shared/store/pgsql"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepositoryReader interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error)
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
}

type UserRepositoryWriter interface {
	CreateUser(ctx context.Context, user *models.Users) (*models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) (*models.Users, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type postgreUserRepository struct {
	db pgsql.Client
}

func NewPostgreRepository(db pgsql.Client) (UserRepositoryReader, UserRepositoryWriter) {
	repo := &postgreUserRepository{db: db}
	return repo, repo
}

func (r *postgreUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error) {
	user := &models.Users{}
	err := r.db.Pool().QueryRow(ctx, getUserByID, id).
		Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.Fullname,
			&user.Status,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
			&user.DeletedAt,
			&user.DeletedBy,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *postgreUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := &models.Users{}
	err := r.db.Pool().QueryRow(ctx, getUserByEmail, email).
		Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.Fullname,
			&user.Status,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
			&user.DeletedAt,
			&user.DeletedBy,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *postgreUserRepository) CreateUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	err := r.db.Pool().QueryRow(ctx, createUser,
		user.Email,
		user.PasswordHash,
		user.Fullname,
		user.Status,
		user.CreatedBy,
		user.UpdatedBy).
		Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *postgreUserRepository) UpdateUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	err := r.db.Pool().QueryRow(ctx, updateUser,
		user.ID,
		user.UpdatedBy,
		user.Email,
		user.PasswordHash,
		user.Fullname,
		user.Status).
		Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *postgreUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Pool().Exec(ctx, deleteUser, id, uuid.Nil)
	return err
}
