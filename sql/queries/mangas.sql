-- name: GetMangaById :one
SELECT * FROM mangas WHERE id = $1;
--

-- name: GetUserFavorites :many
SELECT mangas.*
FROM mangas
JOIN user_mangas ON mangas.id = user_mangas.manga_id
WHERE user_mangas.user_id = $1;

--

-- name: UpdateStatusRead :exec
UPDATE mangas SET read = $1, updated_at = $2
WHERE id = $3;