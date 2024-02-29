package main

import (
	"cipo_cite_server/cmd/internal/config"
	"cipo_cite_server/cmd/internal/storage/postgres"
	testdb "cipo_cite_server/cmd/internal/testDb"

	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	fmt.Printf("envs: %+v\n", cfg.Envs)
	fmt.Printf("config: %+v\n", cfg.Config)

	// init logger: slog
	log := setupLogger(cfg.Config.Env)
	log.Info("starting server on env: " + cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	postgresStore := postgres.PostgresStore{}
	var databaseUrl = fmt.Sprint(
		"postgresql://",
		cfg.Envs.DB_USERNAME,
		":", cfg.Envs.DB_PASSWORD,
		"@", cfg.Envs.DB_HOST, ":", fmt.Sprint(cfg.Envs.DB_PORT), "/", cfg.Envs.DB_DATABASE)
	dbx, err := postgresStore.New(databaseUrl)
	dbx.DB.Ping()
	if err != nil {
		log.Error("cannot connect to database: " + err.Error())
		panic(err)
	}
	log.Info("DB connected: " + cfg.Envs.DB_DATABASE)

	testdb.Testdb(dbx, log)

	// init router: chi, chi render
	// init server:

	// TODO: graceful shutdown
	postgresStore.CloseDB()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log
}
