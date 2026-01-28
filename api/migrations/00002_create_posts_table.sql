-- +goose Up
CREATE TABLE posts (
  id TEXT PRIMARY KEY,
  author_id TEXT NOT NULL,
  body TEXT,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT,
  FOREIGN KEY (author_id) REFERENCES users (id)
);

-- +goose Down
DROP TABLE IF EXISTS posts;