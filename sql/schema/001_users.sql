-- +goose Up
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL,
    api_key TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;