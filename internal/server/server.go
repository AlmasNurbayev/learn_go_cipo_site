package server

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	"cipo_cite_server/internal/storage/postgres"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "net/http/pprof"

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
	//Validator  *validator.Validate
}

func New(version string) *Server {
	return &Server{
		Version: version,
	}
}

func (s *Server) Init() {
	fmt.Println("============================")
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
	//s.Validator = validator.New()
	s.Mux = chi.NewRouter()
	s.HttpServer = &http.Server{
		Addr:         s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port),
		Handler:      s.Mux,
		ReadTimeout:  s.Cfg.Server.Timeout,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  s.Cfg.Server.Idle_timeout,
	}
	s.Mux.Use(middleware.RequestID)
	if s.Cfg.Config.Env == "dev" {
		s.Mux.Use(middleware.Logger)
	}
	s.Mux.Use(middleware.Recoverer)
	s.Mux.Use(middleware.Heartbeat("/ping"))
	// s.Mux.HandleFunc("/debug/pprof/", pprof.Index)
	// s.Mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// s.Mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// s.Mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// s.Mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// s.Mux.HandleFunc("/debug/pprof/goroutine", func(w http.ResponseWriter, r *http.Request) {
	// 	pprof.Handler("goroutine").ServeHTTP(w, r)
	// })
	// //s.Mux.HandleFunc("/threadcreate", pprof.Handler("threadcreate"))
	// //s.Mux.HandleFunc("/heap", pprof.Handler("heap"))
	// s.Mux.HandleFunc("/debug/pprof/block", func(w http.ResponseWriter, r *http.Request) {
	// 	pprof.Handler("block").ServeHTTP(w, r)
	// })
	// s.Mux.HandleFunc("/debug/pprof/mutex", func(w http.ResponseWriter, r *http.Request) {
	// 	pprof.Handler("mutex").ServeHTTP(w, r)
	// })

	s.registerNews()
	s.registerStores()
	s.registerProductsFilters()
	s.registerProduct()

}

func (s *Server) Run() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil {
			s.Log.Error("failed to start server")
		}
	}()
	s.Log.Info("starting Chi server on " + s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port))
	<-done
	s.Log.Warn("stopping Chi server on " + s.Cfg.Server.Addr + ":" + strconv.Itoa(s.Cfg.Server.Http_port))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Cfg.Server.Idle_timeout))

	defer cancel()

	s.Shutdown(ctx)
}

func (s *Server) Shutdown(ctx context.Context) {
	err := s.HttpServer.Shutdown(ctx)
	if err != nil {
		s.Log.Error("server shutdown error: " + err.Error())
	}
	s.Log.Info("server shutdown: " + s.Cfg.Envs.DB_DATABASE)
	s.Sqlx.Close()
	s.Log.Info("DB shutdown: " + s.Cfg.Envs.DB_DATABASE)
}
