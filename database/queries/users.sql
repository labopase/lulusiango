-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    fullname,
    status,
    created_by,
    updated_by
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND deleted_at IS NULL LIMIT 1;

-- name: UpdateUserStatus :one
UPDATE users
SET 
    status = $2,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: UpdateUserProfile :one
UPDATE users
SET 
    fullname = $2,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET 
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $2
WHERE id = $1 AND deleted_at IS NULL;
