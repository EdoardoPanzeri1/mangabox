-- name: GetComicByID :one
SELECT * FROM comics WHERE id = $1;
--

-- name: GetUserFavorites :many
SELECT comics.*
FROM comics
JOIN user_comics ON comics.id = user_comics.comic_id
WHERE user_comics.user_id = $1;

--

-- name: UpdateStatusRead :exec
UPDATE comics SET read = $1, updated_at = $2
WHERE id = $3;