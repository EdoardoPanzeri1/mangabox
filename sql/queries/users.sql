-- name: CreateUser :exec
INSERT INTO users (username, email, password_hash)
VALUES ($1, $2, $3);
--

-- name: UpdateUserDetails :exec
UPDATE users
SET email = $1, password_hash = $2
WHERE username = $3;

-- name: FetchUserByUsername :one
SELECT id, username, email, password_hash
FROM users
WHERE username = $1;

-- name: ListUsers :many
SELECT id, username, email
FROM users;
