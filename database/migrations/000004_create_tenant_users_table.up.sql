CREATE TABLE tenant_users (
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by VARCHAR(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_by VARCHAR(255),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(255),
    PRIMARY KEY (tenant_id, user_id)
);

CREATE INDEX idx_tenant_users_user_id ON tenant_users(user_id) WHERE deleted_at IS NULL;
