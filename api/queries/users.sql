-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- -- name: GetUserByUsername :one
-- SELECT * FROM users WHERE username = ? LIMIT 1;

-- -- name: GetUserByEmail :one
-- SELECT * FROM users WHERE email = ? LIMIT 1;

-- -- name: ListUsers :many
-- SELECT * FROM users ORDER BY created_at DESC;

-- -- name: CreateUser :execresult
-- INSERT INTO users (id, email, name, username, bio, password, avatar_url)
-- VALUES (?, ?, ?, ?, ?, ?, ?);

-- -- name: UpdateUser :exec
-- UPDATE users SET
--   name = COALESCE(?, name),
--   username = COALESCE(?, username),
--   bio = COALESCE(?, bio),
--   password = COALESCE(?, password),
--   avatar_url = COALESCE(?, avatar_url),
--   updated_at = CURRENT_TIMESTAMP
-- WHERE id = ?;

-- -- name: DeleteUser :exec
-- DELETE FROM users WHERE id = ?;
