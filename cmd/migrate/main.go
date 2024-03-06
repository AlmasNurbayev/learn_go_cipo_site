package main

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage"
	"cipo_cite_server/internal/storage/postgres"
	"os"
)

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	// init logger: slog
	log := logger.InitLogger(cfg.Config.Env)
	log.Info("starting migrate on env: " + cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	var postgresStore = postgres.NewStore()
	postgresStore, err := postgresStore.Init(cfg.Envs, log)
	if err != nil {
		log.Error("Error init postgresStore: ", err)
		os.Exit(1)
	}

	// make migrations
	storage.MigrationsUp(postgresStore.Dbx)

	// TODO: graceful shutdown
	postgresStore.Dbx.Close()
	log.Info("DB shutdown: " + cfg.Envs.DB_DATABASE)
}
