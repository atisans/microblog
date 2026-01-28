-- name: GetUserByUsername :one
SELECT id, email, name, username, bio, password, avatar_url, created_at, updated_at FROM users WHERE username = ? LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (id, email, name, username, bio, password, avatar_url, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);
