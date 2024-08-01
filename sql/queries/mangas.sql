-- name: UpdateStatusRead :exec
UPDATE mangas 
SET status = 'read'
WHERE id = $1 AND user_id = $2;

-- name: InsertMangaIntoCatalog :exec
INSERT INTO mangas (
    id, status, user_id, title, issue_number,
    publication_date, storyline, cover_art_url, updated_at,
    images, authors, serializations, genres, explicit_genres,
    themes, demographics, score, scored_by, rank,
    popularity, members, favorites, synopsis, background,
    relations, external_links
)
VALUES (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9,
    $10, $11, $12, $13, $14,
    $15, $16, $17, $18, $19,
    $20, $21, $22, $23, $24,
    $25, $26
);

-- name: RetrieveCatalog :many
SELECT m.title, m.authors, m.status, m.cover_art_url, m.issue_number
FROM mangas m
JOIN users u ON m.user_id = u.id
WHERE m.user_id = $1;

-- name: DeleteManga :exec
DELETE FROM mangas
WHERE id = $1 AND user_id = $2;