package newsHandler

import (
	"cipo_cite_server/internal/models"
	"cipo_cite_server/internal/repository/news"
	"cipo_cite_server/internal/server/handlers/news/newsFilter"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Operations interface {
	List() (*[]models.News, error)
}

func NewsGetAll(w http.ResponseWriter, r *http.Request, repo *news.RepositoryDb, log *slog.Logger) {
	//params := r.URL.Query()
	filters := newsFilter.Filters(r.URL.Query())

	value, err := repo.List(filters)
	if err != nil {
		log.Error("error on NewsGetAll: ", err)
		http.Error(w, "Ошибка получения из БД", http.StatusInternalServerError)
		return
	}
	valueJSON, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "Ошибка преобразования JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(valueJSON)
}
