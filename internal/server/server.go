package server

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage/postgres"
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Version    string
	Cfg        *config.MultiConfig
	Sqlx       *sqlx.DB
	Log        *slog.Logger
	Mux        *chi.Mux
	HttpServer *http.Server
	Addr       string
}

func New(version string) *Server {
	return &Server{
		Version: version,
	}
}

func (s *Server) Init() {
	s.Cfg = config.MustLoad()
	s.Log = logger.InitLogger(s.Cfg.Config.Env)
	s.Log.Info("init server on env: " + s.Cfg.Config.Env)
	s.Log.Debug("debug message is enabled")
	var postgresStore = postgres.NewStore()
	postgresStore, err := postgresStore.Init(s.Cfg.Envs, s.Log)
	if err != nil {
		s.Log.Error("Error init postgresStore:", err)
		panic(err)
	}
	s.Sqlx = postgresStore.Dbx
	s.Mux = chi.NewRouter()
	s.HttpServer = &http.Server{
		Addr:         s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port),
		Handler:      s.Mux,
		ReadTimeout:  time.Duration(s.Cfg.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(s.Cfg.Server.Timeout) * time.Second,
		IdleTimeout:  time.Duration(s.Cfg.Server.Idle_timeout) * time.Second,
	}
}

func (s *Server) Run() {
	s.Log.Info("starting server on " + s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port))
	if err := s.HttpServer.ListenAndServe(); err != nil {
		s.Log.Error("failed to start server")
	}

	s.Shutdown()
}

func (s *Server) Shutdown() {
	s.HttpServer.Shutdown(context.Background())
	s.Sqlx.Close()
	s.Log.Info("DB shutdown: " + s.Cfg.Envs.DB_DATABASE)
}
