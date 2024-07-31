-- +goose Up
ALTER TABLE mangas ADD COLUMN status status DEFAULT 'bought';

-- +goose Down
ALTER TABLE mangas DROP COLUMN status;