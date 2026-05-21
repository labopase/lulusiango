package models

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusPending   UserStatus = "pending"
	UserStatusSuspended UserStatus = "suspended"
)

func (e *UserStatus) Validate(v UserStatus) error {
	if !slices.Contains([]UserStatus{UserStatusActive, UserStatusPending, UserStatusSuspended}, v) {
		return fmt.Errorf("invalid user status: %s", v)
	}

	return nil
}

type Users struct {
	ID           uuid.UUID        `json:"id"`
	Email        string           `json:"email"`
	PasswordHash string           `json:"password_hash"`
	Fullname     string           `json:"fullname"`
	Status       UserStatus       `json:"status"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	CreatedBy    *string          `json:"created_by"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	UpdatedBy    *string          `json:"updated_by"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
	DeletedBy    *string          `json:"deleted_by"`
}
