package main

import (
	"database/sql"
	"log/slog"
	"os"
	"runtime/debug"
	"strconv"
	"sync"

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
	baseURL  string
	httpPort int
	db       struct {
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

	if v := os.Getenv("BASE_URL"); v != "" {
		cfg.baseURL = v
	} else {
		cfg.baseURL = "http://localhost:6364"
	}

	if v := os.Getenv("HTTP_PORT"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			cfg.httpPort = i
		}
	} else {
		cfg.httpPort = 6364
	}

	if v := os.Getenv("DATABASE_URL"); v != "" {
		cfg.db.dsn = v
	} else {
		cfg.db.dsn = "microblog.db?_foreign_keys=on"
	}

	db, err := database.NewDB(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := database.New(db)

	app := &application{
		config:  cfg,
		db:      db,
		queries: queries,
		logger:  logger,
	}

	return app.serveHTTP()
}
