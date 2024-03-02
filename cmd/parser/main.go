package main

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/parser"
	"cipo_cite_server/internal/storage/postgres"
)

func main() {

	// init config: cleanenv
	Cfg := config.MustLoad()

	// init logger: slog
	log := logger.InitLogger(Cfg.Config.Env)
	log.Info("starting parser on env: " + Cfg.Config.Env)
	log.Debug("debug message is enabled")

	// init storage: pgx, sqlx
	db := postgres.InitPostgresStore(Cfg.Envs, log)

	// init moved files
	// result, err := parser.MovedInputFiles(Cfg.Config, log)
	// if err != nil {
	// 	log.Error("Error moving input files:", err)
	// 	os.Exit(1)
	// }
	// if parser.MovedImages(result.NewPath, log) != nil {
	// 	log.Error("Error moving images:", err)
	// 	os.Exit(1)
	// }

	//временный массив для отладки
	result := parser.MovedInputFilesT{
		Files: []parser.InputFilesT{
			{TypeFile: "classificator", PathFile: "input/import0_1.xml"},
			{TypeFile: "offer", PathFile: "input/offers0_1.xml"},
		},
		NewPath: "newPath",
	}

	parser.ReadAndParse(&result, log, db)

	// init parser

	// TODO: graceful shutdown
	db.Close()
	log.Info("DB shutdown: " + Cfg.Envs.DB_DATABASE)
}
