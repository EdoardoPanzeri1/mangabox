-- +goose Up
CREATE TYPE status AS ENUM ('bought', 'reading', 'completed');

-- +goose Down
DROP TYPE status;