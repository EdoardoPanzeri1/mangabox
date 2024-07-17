-- +goose Up
CREATE TABLE comics (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    issue_number INTEGER NOT NULL,
    publication_date DATE NOT NULL,
    storyline TEXT,
    cover_art_url TEXT,
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE comics;