package main

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage/postgres"
	"embed"

	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	// init logger: slog
	log := logger.InitLogger(cfg.Config.Env)
	log.Info("starting migrate on env: " + cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	db := postgres.InitPostgresStore(cfg.Envs, log)

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, cfg.Db.Migrations_path); err != nil {
		panic(err)
	}

	// init moved files
	// init parser

	// TODO: graceful shutdown
	db.Close()
	log.Info("DB shutdown: " + cfg.Envs.DB_DATABASE)
}
