CREATE TYPE tenant_status AS ENUM ('active', 'suspended', 'archived');
CREATE TYPE user_status AS ENUM ('active', 'pending', 'suspended');
CREATE TYPE api_key_status AS ENUM ('active', 'rotated', 'revoked', 'expired');
