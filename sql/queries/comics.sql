-- name: GetComicByID :one
SELECT * FROM comics WHERE id = ?;
--

-- name: GetUserFavorites: many
SELECT * FROM comics WHERE id IN (SELECT comic_id FROM user_comics WHERE user_id = ?)
--

-- name: UpdateStatusRead :exec
UPDATE comics SET read = ?, updated_at = ?
WHERE id = ?