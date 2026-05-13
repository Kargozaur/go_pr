-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),

    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles (
    id UUID PRIMARY KEY DEFAULT uuidv7(),

    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,

    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT uuidv7(),

    token_hash VARCHAR(150) NOT NULL,

    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_refresh_tokens_user_id
    ON refresh_tokens(user_id);

CREATE INDEX idx_refresh_tokens_deleted_at
    ON refresh_tokens(deleted_at);

-- +goose Down
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS users;
