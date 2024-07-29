-- +goose Up
CREATE TYPE status AS ENUM ('bought', 'read');

-- +goose Down
DROP TYPE IF EXISTS status;