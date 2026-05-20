CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    key_id UUID NOT NULL,
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    name VARCHAR(100) NOT NULL,
    key_hash VARCHAR(255) UNIQUE NOT NULL,
    status api_key_status NOT NULL DEFAULT 'active',
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(255)
);

CREATE INDEX idx_api_keys_key_hash ON api_keys(key_hash) WHERE deleted_at IS NULL;
CREATE INDEX idx_api_keys_key_id ON api_keys(key_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_api_keys_tenant_user ON api_keys(tenant_id, user_id) WHERE deleted_at IS NULL;
