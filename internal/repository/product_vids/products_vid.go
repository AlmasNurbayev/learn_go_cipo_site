package product_vids

import (
	"cipo_cite_server/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(product_group models.ProductVids) (int64, error)
	Update(product_group models.ProductVids) (int64, error)
	List() (*[]models.ProductVids, error)
}

type repository struct {
	db *sqlx.Tx
}

func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(product_group models.ProductVids) (int64, error) {
	query := `INSERT INTO product_vids
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
		//utils.PrintAsJSON(res)
	}
	return res, nil
}

func (s *repository) Update(product_group models.ProductVids) (int64, error) {
	if product_group.Id == 0 {
		return 0, errors.New("id is empty")
	}
	query := `UPDATE product_vids
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

func (s *repository) List() (*[]models.ProductVids, error) {
	query := `SELECT * FROM product_vids`
	var res []models.ProductVids
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
