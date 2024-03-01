package main

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/parser"
	"cipo_cite_server/internal/storage/postgres"
)

func main() {

	// init config: cleanenv
	cfg := config.MustLoad()

	// init logger: slog
	log := logger.InitLogger(cfg.Config.Env)
	log.Info("starting parser on env: " + cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	db := postgres.InitPostgresStore(cfg.Envs, log)

	// init moved files
	//files := parser.MovedInputFiles(cfg.Config, log)
	// временный массив
	files := []parser.InputFilesT{
		{TypeFile: "classificator", PathFile: "input/import0_1.xml"},
		{TypeFile: "offer", PathFile: "input/offers0_1.xml"},
	}

	parser.Parser(&files, log)

	// init parser

	// TODO: graceful shutdown
	db.Close()
	log.Info("DB shutdown: " + cfg.Envs.DB_DATABASE)
}
