package productHandler

import (
	"cipo_cite_server/internal/repository/products"
	"cipo_cite_server/internal/repository/qnt_price_registry"
	"cipo_cite_server/internal/server/handlers/product/productFilter"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
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
	_, err = w.Write(outputJSON)
	if err != nil {
		log.Error("Error Product GetById write JSON")
		http.Error(w, "Error Product GetById write JSON", http.StatusInternalServerError)
		return
	}
}

func ProductNews(w http.ResponseWriter, r *http.Request, log *slog.Logger,
	qntPriceRepo *qnt_price_registry.RepositoryDb,
	productRepo *products.RepositoryDb,
) {

	params := r.URL.Query()
	var newsCount int
	var err error
	if params.Get("news") == "" {
		http.Error(w, "not content parameter news", http.StatusBadRequest)
		return
	} else {
		newsString := params.Get("news")
		newsCount, err = strconv.Atoi(newsString)
		if err != nil {
			http.Error(w, "not correct parameter news", http.StatusBadRequest)
			return
		}
	}

	lastRegistrator, err := qntPriceRepo.GetLastRegistratorsFromQntPrices()
	if err != nil {
		log.Error("error on productNews - find last registrator: ", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	products, err := productRepo.GetProductNews(newsCount, lastRegistrator)
	if err != nil {
		log.Error("error on product productNews - DB: ", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	outputJSON, err := json.Marshal(products)
	if err != nil {
		log.Error("error ProductNews  marshal JSON: ", err)
		http.Error(w, "Error ProductNews marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(outputJSON)
	if err != nil {
		log.Error("Error ProductNews write JSON")
		http.Error(w, "Error ProductNews write JSON", http.StatusInternalServerError)
		return
	}
}

func List(w http.ResponseWriter, r *http.Request, log *slog.Logger,
	qntPriceRepo *qnt_price_registry.RepositoryDb,
	productRepo *products.RepositoryDb,
) {
	params := r.URL.Query()

	filters, err := productFilter.FiltersList(params)
	if err != nil {
		log.Error("error on products - parsing filters: ", err)
		http.Error(w, "not correct filters", http.StatusBadRequest)
		return
	}
	lastRegistrator, err := qntPriceRepo.GetLastRegistratorsFromQntPrices()
	if err != nil {
		log.Error("error on products - find last registrator: ", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	products, err := productRepo.List(*filters, lastRegistrator)
	if err != nil {
		log.Error("error on products - DB: ", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	outputJSON, err := json.Marshal(products)
	if err != nil {
		log.Error("error products  marshal JSON: ", err)
		http.Error(w, "Error products marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(outputJSON)
	if err != nil {
		log.Error("Error products write JSON")
		http.Error(w, "Error products write JSON", http.StatusInternalServerError)
		return
	}
}
