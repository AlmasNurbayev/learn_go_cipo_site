package server

import (
	"cipo_cite_server/internal/repository/news"
	newsHandler "cipo_cite_server/internal/server/handlers/news"
	"net/http"
)

func (s *Server) registerNews() {
	newsRepo := news.NewRepositoryDb(s.Sqlx)
	s.Mux.Get("/news", func(w http.ResponseWriter, r *http.Request) {
		newsHandler.NewsGetAll(w, r, newsRepo, s.Log)
	})
}
