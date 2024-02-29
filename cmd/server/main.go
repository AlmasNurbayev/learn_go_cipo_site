package main

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage/postgres"
	testdb "cipo_cite_server/internal/testDb"

	"fmt"
)

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	fmt.Printf("envs: %+v\n", cfg.Envs)
	fmt.Printf("config: %+v\n", cfg.Config)

	// init logger: slog
	log := logger.InitLogger(cfg.Config.Env)
	log.Info("starting server on env: " + cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	dbPostgres := postgres.InitPostgresStore(cfg.Envs, log)
	testdb.Testdb(dbPostgres, log)

	// init router: chi, chi render
	// init server:

	// TODO: graceful shutdown
	dbPostgres.Close()
	log.Info("DB shutdown: " + cfg.Envs.DB_DATABASE)

}
