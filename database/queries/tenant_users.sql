-- name: AddUserToTenant :one
INSERT INTO tenant_users (
    tenant_id,
    user_id,
    created_by,
    updated_by
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT (tenant_id, user_id) 
DO UPDATE SET 
    deleted_at = NULL,
    deleted_by = NULL,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = EXCLUDED.updated_by
RETURNING *;

-- name: RemoveUserFromTenant :one
UPDATE tenant_users
SET
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $3
WHERE tenant_id = $1 AND user_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: GetTenantUser :one
SELECT * FROM tenant_users
WHERE tenant_id = $1 AND user_id = $2 AND deleted_at IS NULL;

-- name: ListUsersInTenant :many
SELECT u.* FROM users u
INNER JOIN tenant_users tu ON tu.user_id = u.id
WHERE tu.tenant_id = $1 AND tu.deleted_at IS NULL AND u.deleted_at IS NULL
ORDER BY tu.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListTenantsForUser :many
SELECT t.* FROM tenants t
INNER JOIN tenant_users tu ON tu.tenant_id = t.id
WHERE tu.user_id = $1 AND tu.deleted_at IS NULL AND t.deleted_at IS NULL
ORDER BY tu.created_at DESC
LIMIT $2 OFFSET $3;
