package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"runtime/debug"
	"strconv"
	"sync"

	_ "github.com/joho/godotenv/autoload"

	"microblog/internal/database"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	port      int
	jwtSecret string
	db        struct {
		dsn string
	}
}

type application struct {
	config  config
	db      *sql.DB
	queries *database.Queries
	logger  *slog.Logger
	wg      sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	if v := os.Getenv("PORT"); v == "" {
		if i, err := strconv.Atoi(v); err == nil {
			cfg.port = i
		}
	}

	if v := os.Getenv("DATABASE_URL"); v == "" {
		return errors.New("DATABASE_URL environment variable is not set")
	} else {
		cfg.db.dsn = v
	}

	if v := os.Getenv("JWT_SECRET_KEY"); v == "" {
		return errors.New("JWT_SECRET_KEY environment variable is not set")
	} else {
		cfg.jwtSecret = v
	}

	db, err := database.NewDB(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Apply SQLite pragmas for performance optimization (matching TypeScript config)
	pragmas := []string{
		"PRAGMA journal_mode = WAL",
		"PRAGMA busy_timeout = 5000",
		"PRAGMA synchronous = NORMAL",
		"PRAGMA cache_size = 10000",
		"PRAGMA foreign_keys = ON",
		"PRAGMA temp_store = MEMORY",
		"PRAGMA mmap_size = 268435456",
	}

	for _, pragma := range pragmas {
		if _, err := db.ExecContext(context.Background(), pragma); err != nil {
			return err
		}
	}

	queries := database.New(db)

	app := &application{
		config:  cfg,
		db:      db,
		queries: queries,
		logger:  logger,
		wg:      sync.WaitGroup{},
	}

	return app.serveHTTP()
}
