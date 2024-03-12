package server

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage/postgres"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	validator "github.com/go-playground/validator/v10"
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
	Validator  *validator.Validate
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
	s.Validator = validator.New()
	s.Mux = chi.NewRouter()
	s.HttpServer = &http.Server{
		Addr:         s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port),
		Handler:      s.Mux,
		ReadTimeout:  s.Cfg.Server.Timeout,
		WriteTimeout: s.Cfg.Server.Timeout,
		IdleTimeout:  s.Cfg.Server.Idle_timeout,
	}
	s.Mux.Use(middleware.RequestID)
	s.Mux.Use(middleware.Logger)
	s.Mux.Use(middleware.Recoverer)
	s.Mux.Use(middleware.Heartbeat("/ping"))
	s.registerNews()
	s.registerStores()
	s.registerProductsFilters()

}

func (s *Server) Run() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil {
			s.Log.Error("failed to start server")
		}
	}()
	s.Log.Info("starting server on " + s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port))
	<-done
	s.Log.Warn("stopping server on " + s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Cfg.Server.Idle_timeout))

	defer cancel()

	s.Shutdown(ctx)
}

func (s *Server) Shutdown(ctx context.Context) {
	s.HttpServer.Shutdown(ctx)
	s.Log.Info("server shutdown: " + s.Cfg.Envs.DB_DATABASE)
	s.Sqlx.Close()
	s.Log.Info("DB shutdown: " + s.Cfg.Envs.DB_DATABASE)
}
