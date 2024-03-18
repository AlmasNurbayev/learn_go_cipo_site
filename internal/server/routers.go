package server

import (
	"cipo_cite_server/internal/repository/news"
	"cipo_cite_server/internal/repository/products"
	"cipo_cite_server/internal/repository/products_groups"
	"cipo_cite_server/internal/repository/qnt_price_registry"
	"cipo_cite_server/internal/repository/sizes"
	"cipo_cite_server/internal/repository/stores"
	"cipo_cite_server/internal/repository/vids"
	newsHandler "cipo_cite_server/internal/server/handlers/news"
	productHandler "cipo_cite_server/internal/server/handlers/product"
	productFiltersHandler "cipo_cite_server/internal/server/handlers/productFilters"
	storeHandler "cipo_cite_server/internal/server/handlers/stores"
	"net/http"
)

func (s *Server) registerNews() {
	newsRepo := news.NewRepositoryDb(s.Sqlx)
	s.Mux.Get("/news", func(w http.ResponseWriter, r *http.Request) {
		newsHandler.NewsGetAll(w, r, newsRepo, s.Log)
	})
	s.Mux.Get("/newsID", func(w http.ResponseWriter, r *http.Request) {
		newsHandler.NewsGetID(w, r, newsRepo, s.Log)
	})
}

func (s *Server) registerStores() {
	storesRepo := stores.NewRepositoryDb(s.Sqlx)
	s.Mux.Get("/stores", func(w http.ResponseWriter, r *http.Request) {
		storeHandler.StoresGetAll(w, r, storesRepo, s.Log)
	})
}

func (s *Server) registerProductsFilters() {
	storesRepo := stores.NewRepositoryDb(s.Sqlx)
	sizeRepo := sizes.NewRepositoryDb(s.Sqlx)
	productGroupRepo := products_groups.NewRepositoryDb(s.Sqlx)
	vidsRepo := vids.NewRepositoryDb(s.Sqlx)

	s.Mux.Get("/productFilters", func(w http.ResponseWriter, r *http.Request) {
		productFiltersHandler.GetAll(w, r, s.Log, storesRepo, sizeRepo, productGroupRepo, vidsRepo)
	})
}

func (s *Server) registerProduct() {
	qntPriceRepo := qnt_price_registry.NewRepositoryDb(s.Sqlx)
	productRepo := products.NewRepositoryDb(s.Sqlx)

	s.Mux.Get("/productNews", func(w http.ResponseWriter, r *http.Request) {
		productHandler.ProductNews(w, r, s.Log, qntPriceRepo, productRepo)
	})

	s.Mux.Get("/product", func(w http.ResponseWriter, r *http.Request) {
		productHandler.GetById(w, r, s.Log, qntPriceRepo, productRepo)
	})

	s.Mux.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		productHandler.List(w, r, s.Log, qntPriceRepo, productRepo)
	})

}
