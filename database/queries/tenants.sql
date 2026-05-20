-- name: CreateTenant :one
INSERT INTO tenants (
    name,
    slug,
    status,
    created_by,
    updated_by
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTenantByID :one
SELECT * FROM tenants
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetTenantBySlug :one
SELECT * FROM tenants
WHERE slug = $1 AND deleted_at IS NULL;

-- name: ListTenants :many
SELECT * FROM tenants
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateTenant :one
UPDATE tenants
SET
    name = COALESCE(sqlc.narg('name'), name),
    slug = COALESCE(sqlc.narg('slug'), slug),
    status = COALESCE(sqlc.narg('status'), status),
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteTenant :one
UPDATE tenants
SET
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;
