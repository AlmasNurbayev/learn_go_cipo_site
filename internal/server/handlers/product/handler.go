package productHandler

import (
	"cipo_cite_server/internal/repository/products"
	"cipo_cite_server/internal/repository/qnt_price_registry"
	"cipo_cite_server/internal/server/handlers/product/productFilter"
	"encoding/json"
	"log/slog"
	"net/http"
)

func GetById(w http.ResponseWriter, r *http.Request, log *slog.Logger,
	qntPriceRepo *qnt_price_registry.RepositoryDb,
	productRepo *products.RepositoryDb,
) {

	params := r.URL.Query()

	filters, err := productFilter.Filters(params)
	if err != nil {
		log.Error("error on product GetById - parsing filters: ", err)
		http.Error(w, "not correct filters", http.StatusBadRequest)
		return
	}
	//utils.PrintAsJSON(filters)

	if params.Get("id") == "" && params.Get("name_1c") == "" {
		http.Error(w, "not content id or name_1c", http.StatusBadRequest)
		return
	}

	lastRegistrator, err := qntPriceRepo.GetLastRegistratorsFromQntPrices()
	if err != nil {
		log.Error("error on product GetById - find last registrator: ", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	product, err := productRepo.GetByIdOrName(*filters, lastRegistrator)
	if err != nil {
		log.Error("error on product GetById - DB: ", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	outputJSON, err := json.Marshal(product)
	if err != nil {
		log.Error("error on Product GetById marshal JSON: ", err)
		http.Error(w, "Error Product GetById marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(outputJSON)
}
