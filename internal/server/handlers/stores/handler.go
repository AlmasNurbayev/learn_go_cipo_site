package storesHandler

import (
	"cipo_cite_server/internal/repository/stores"
	"encoding/json"
	"log/slog"
	"net/http"
)

func StoresGetAll(w http.ResponseWriter, r *http.Request, repo *stores.RepositoryDb, log *slog.Logger) {

	value, err := repo.List()
	if err != nil {
		log.Error("error on StoresGetAll: ", err)
		http.Error(w, "Error DB query ", http.StatusInternalServerError)
		return
	}
	valueJSON, err := json.Marshal(value)
	if err != nil {
		log.Error("Error StoresGetAll marshal JSON")
		http.Error(w, "Error StoresGetAll marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(valueJSON)
}
