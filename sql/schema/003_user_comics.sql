-- +goose Up
CREATE TABLE user_comics (
    id INTEGER PRIMARY KEY,
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
    comic_id TEXT REFERENCES comics(id) ON DELETE CASCADE,
    favorite BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE user_comics;