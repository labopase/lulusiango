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
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND deleted_at IS NULL;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE(sqlc.narg('email'), email),
    password_hash = COALESCE(sqlc.narg('password_hash'), password_hash),
    fullname = COALESCE(sqlc.narg('fullname'), fullname),
    status = COALESCE(sqlc.narg('status'), status),
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteUser :one
UPDATE users
SET
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;
