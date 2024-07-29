-- +goose Up
CREATE TYPE status AS ENUM ('bought', 'read');

-- +goose Down
ALTER TABLE mangas DROP COLUMN status;

DROP TYPE IF EXISTS status;