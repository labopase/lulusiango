-- name: CreateAPIKey :one
INSERT INTO api_keys (
    key_id,
    tenant_id,
    user_id,
    name,
    key_hash,
    status,
    expires_at,
    created_by,
    updated_by
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetAPIKeyByHash :one
SELECT * FROM api_keys
WHERE key_hash = $1 AND deleted_at IS NULL AND (expires_at IS NULL OR expires_at > CURRENT_TIMESTAMP);

-- name: GetAPIKeyByID :one
SELECT * FROM api_keys
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListAPIKeys :many
SELECT * FROM api_keys
WHERE tenant_id = $1 AND user_id = $2 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;

-- name: UpdateAPIKeyStatus :one
UPDATE api_keys
SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: RevokeAPIKey :one
UPDATE api_keys
SET
    status = 'revoked',
    deleted_at = CURRENT_TIMESTAMP,
    deleted_by = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;
