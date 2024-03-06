package vids

import (
	"cipo_cite_server/internal/models"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(product_group models.ProductsGroup) (int64, error)
	List() (*[]models.ProductsGroup, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(product_group models.VidsModel) (int64, error) {
	query := `INSERT INTO vids 
	(id_1c, name_1c, registrator_id) 
		VALUES 
		(:id_1c, :name_1c, :registrator_id) 
		RETURNING id`

	rows, err := s.db.NamedQuery(query, product_group)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var res int64

	for rows.Next() {
		err := rows.Scan(&res)
		if err != nil {
			return 0, err
		}
	}
	return res, nil
}

func (s *repository) List() (*[]models.VidsModel, error) {
	query := `SELECT * FROM vids`
	var res []models.VidsModel
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
