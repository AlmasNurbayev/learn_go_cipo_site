package product_desc

import (
	"cipo_cite_server/internal/models"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	List() (*[]models.ProductsDesc, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) List() (*[]models.ProductsDesc, error) {
	query := `SELECT * FROM product_desc`
	var res []models.ProductsDesc
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
