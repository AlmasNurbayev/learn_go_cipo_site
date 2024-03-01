package main

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage"
	"cipo_cite_server/internal/storage/postgres"
)

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	// init logger: slog
	log := logger.InitLogger(cfg.Config.Env)
	log.Info("starting migrate on env: " + cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	db := postgres.InitPostgresStore(cfg.Envs, log)

	storage.MigrationsUp(db)

	// init moved files
	// init parser

	// TODO: graceful shutdown
	db.Close()
	log.Info("DB shutdown: " + cfg.Envs.DB_DATABASE)
}
