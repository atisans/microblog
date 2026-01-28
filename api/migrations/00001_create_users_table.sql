-- +goose Up
CREATE TABLE users (
  id TEXT PRIMARY KEY,
  email TEXT UNIQUE,
  name TEXT,
  username TEXT UNIQUE NOT NULL,
  bio TEXT,
  password TEXT,
  avatar_url TEXT,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT
);

-- +goose Down
DROP TABLE IF EXISTS users;