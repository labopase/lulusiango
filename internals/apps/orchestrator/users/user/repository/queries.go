package repository

const createUser = `
INSERT INTO users (
    email,
    password_hash,
    fullname,
    status,
    created_by,
    updated_by
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id
`

const deleteUser = `
UPDATE users
SET
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $2
WHERE id = $1 AND deleted_at IS NULL
`
const getUserByEmail = `
SELECT id, email, password_hash, fullname, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM users
WHERE email = $1 AND deleted_at IS NULL
`

const getUserByID = `
SELECT id, email, password_hash, fullname, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM users
WHERE id = $1 AND deleted_at IS NULL
`

const updateUser = `
UPDATE users
SET
    email = COALESCE($3, email),
    password_hash = COALESCE($4, password_hash),
    fullname = COALESCE($5, fullname),
    status = COALESCE($6, status),
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $2
WHERE id = $1 AND deleted_at IS NULL
`
