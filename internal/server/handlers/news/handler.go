package newsHandler

import (
	"cipo_cite_server/internal/repository/news"
	"cipo_cite_server/internal/server/handlers/news/newsFilter"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

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

func NewsGetID(w http.ResponseWriter, r *http.Request, repo *news.RepositoryDb, log *slog.Logger) {
	params := r.URL.Query()
	idString := params.Get("id")
	if idString == "" {
		http.Error(w, "not content id", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "id not correct", http.StatusBadRequest)
		return
	}

	value, err := repo.ListById(int64(idInt))
	if err != nil {
		log.Error("error on NewsGetById: ", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if len(*value) == 0 {
		http.Error(w, "not found id", http.StatusBadRequest)
		return
	}
	valueJSON, err := json.Marshal((*value)[0]) // возвращаем не массив, а сам объект
	if err != nil {
		log.Error("error on NewsGetById: ", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(valueJSON)
}
