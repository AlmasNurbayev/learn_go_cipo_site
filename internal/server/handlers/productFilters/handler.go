package productFiltersHandler

import (
	"cipo_cite_server/internal/models"
	"cipo_cite_server/internal/repository/products_groups"
	"cipo_cite_server/internal/repository/sizes"
	"cipo_cite_server/internal/repository/stores"
	"cipo_cite_server/internal/repository/vids"
	"encoding/json"
	"log/slog"
	"net/http"
)

type output struct {
	Stores        []models.StoresShort        `json:"store"`
	Sizes         []models.SizesShort         `json:"size"`
	ProductGroups []models.ProductsGroupShort `json:"product_group"`
	Vids          []models.VidsShort          `json:"vid_modeli"`
}

func GetAll(w http.ResponseWriter, r *http.Request, log *slog.Logger,
	storesRepo *stores.RepositoryDb, sizeRepo *sizes.RepositoryDb,
	productGroupRepo *products_groups.RepositoryDb, vidsRepo *vids.RepositoryDb) {

	stores, err := storesRepo.ListShort()
	if err != nil {
		log.Error("error on ProductFilters: ", err)
		http.Error(w, "Error DB query ", http.StatusInternalServerError)
		return
	}
	sizes, err := sizeRepo.ListShort()
	if err != nil {
		log.Error("error on ProductFilters: ", err)
		http.Error(w, "Error DB query ", http.StatusInternalServerError)
		return
	}
	vids, err := vidsRepo.ListShort()
	if err != nil {
		log.Error("error on ProductFilters: ", err)
		http.Error(w, "Error DB query ", http.StatusInternalServerError)
		return
	}
	productGroups, err := productGroupRepo.ListShort()
	if err != nil {
		log.Error("error on ProductFilters: ", err)
		http.Error(w, "Error DB query ", http.StatusInternalServerError)
		return
	}

	output := output{
		Stores:        *stores,
		Sizes:         *sizes,
		ProductGroups: *productGroups,
		Vids:          *vids,
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Error("Error ProductFilters marshal JSON")
		http.Error(w, "Error ProductFilters marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(outputJSON)
	if err != nil {
		log.Error("Error ProductFilters write JSON")
		http.Error(w, "Error ProductFilters write JSON", http.StatusInternalServerError)
		return
	}
}
