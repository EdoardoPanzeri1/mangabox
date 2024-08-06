-- +goose Up
CREATE TABLE IF NOT EXISTS mangas (
    id TEXT PRIMARY KEY,
    status status DEFAULT 'bought',  -- Use the ENUM type created earlier
    user_id UUID REFERENCES users(id),
    title TEXT NOT NULL,
    issue_number INTEGER NOT NULL,
    publication_date DATE NOT NULL,
    storyline TEXT,
    cover_art_url TEXT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- New fields
    images JSON,                       -- Store as JSON
    authors JSON,                      -- Store as JSON
    serializations JSON,               -- Store as JSON
    genres JSON,                       -- Store as JSON
    explicit_genres JSON,              -- Store as JSON
    themes JSON,                       -- Store as JSON
    demographics JSON,                 -- Store as JSON
    score DOUBLE PRECISION,            -- Store float as DOUBLE
    scored_by INTEGER,                 -- Store int
    rank INTEGER,                      -- Store int
    popularity INTEGER,                -- Store int
    members INTEGER,                   -- Store int
    favorites INTEGER,                 -- Store int
    synopsis TEXT,
    background TEXT,
    relations JSON,                    -- Store as JSON
    external_links JSON                -- Store as JSON
);

ALTER TABLE mangas
ALTER COLUMN user_id TYPE UUID USING user_id::UUID;

-- +goose Down
DROP TABLE IF EXISTS mangas;