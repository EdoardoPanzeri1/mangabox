-- +goose Up
CREATE TABLE user_mangas (
    id INTEGER PRIMARY KEY,
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
    manga_id TEXT REFERENCES mangas(id) ON DELETE CASCADE,
    favorite BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE user_mangas;