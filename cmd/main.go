package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// load config
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: os.Getenv("DATABASE_URL"),
		},
	}

	// connect to db
	pool, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Pool unable to be created", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("Database unable to be pinged!", "error", err)
		os.Exit(1)
	}

	logger.Info("connected to database pool", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     pool,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
}
