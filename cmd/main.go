package main

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"pcr/internal/handler"
)

type config struct {
	DbUrl     string `env:"DB_URL"`
	SecretKey string `env:"SECRET_KEY"`
	RootDir   string `env:"ROOT_DIR" envDefault:"/var/pcr"`
	Listen    string `env:"LISTEN" envDefault:":8100"`
}

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	cfg, err := env.ParseAs[config]()
	if err != nil {
		return 1
	}

	err = start(ctx, cfg)
	if err != nil {
		return 1
	}

	return 0
}

func start(ctx context.Context, cfg config) error {
	u, err := url.Parse(cfg.DbUrl)
	if err != nil {
		slog.Error("Invalid database URL", "reason", err)
		return err
	}
	slog.Info(
		"Connecting to database",
		slog.String("host", u.Host),
		slog.String("port", u.Port()),
		slog.String("user", u.User.Username()),
		slog.String("database", u.Path[1:]),
	)

	pool, err := pgxpool.New(ctx, cfg.DbUrl)
	if err != nil {
		slog.Error("Failed to establish database connection", "reason", err)
	}

	pcr := &handler.Pcr{
		Cursor: pool,
		Root:   cfg.RootDir,
	}
	http.Handle("GET /{tenant}/{file}", pcr)

	slog.Info("Starting HTTP server", "listen", cfg.Listen)
	return http.ListenAndServe(cfg.Listen, nil)
}
