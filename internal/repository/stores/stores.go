package stores

import (
	"cipo_cite_server/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(product_group models.ProductsGroup) (int64, error)
	Update(product_group models.ProductsGroup) (int64, error)
	List() (*[]models.ProductsGroup, error)
}

type repository struct {
	db *sqlx.Tx
}

func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(product_group models.Stores) (int64, error) {

	// записываем только часть полей - остальные для правки вручную
	query := `INSERT INTO stores 
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

func (s *repository) Update(product_group models.Stores) (int64, error) {
	if product_group.Id == 0 {
		return 0, errors.New("id is empty")
	}

	// записываем только часть полей - остальные для правки вручную
	query := `UPDATE stores
	SET id_1c = :id_1c, name_1c = :name_1c, registrator_id = :registrator_id, changed_at = CURRENT_TIMESTAMP
	WHERE id = :id RETURNING id`

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

func (s *repository) List() (*[]models.Stores, error) {
	query := `SELECT * FROM stores`
	var res []models.Stores
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
