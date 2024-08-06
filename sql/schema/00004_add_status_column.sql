-- +goose Up
ALTER TABLE mangas ADD COLUMN IF NOT EXISTS status status DEFAULT 'bought';

-- +goose Down
ALTER TABLE mangas DROP COLUMN status;