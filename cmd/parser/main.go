package main

import (
	"cipo_cite_server/internal/parser"
)

var Version = "v0.1.0"

func main() {

	p := parser.New(Version)
	p.Init()
	p.Run()

	// init config: cleanenv
	// Cfg := config.MustLoad()

	// init logger: slog
	// log := logger.InitLogger(Cfg.Config.Env)
	// log.Info("starting parser on env: " + Cfg.Config.Env)
	// log.Debug("debug message is enabled")

	// // init storage: pgx, sqlx
	// var postgresStore = postgres.NewStore()
	// postgresStore, err := postgresStore.Init(Cfg.Envs, log)
	// if err != nil {
	// 	log.Error("Error init postgresStore:", err)
	// 	os.Exit(1)
	// }

	//временный массив для отладки

}
