package priceVids

import (
	"cipo_cite_server/internal/models"
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(size models.Sizes) (int64, error)
	Update(size models.Sizes) (int64, error)
	List() (*[]models.Sizes, error)
}

type repository struct {
	db *sqlx.Tx
}

func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(size models.PriceVids) (int64, error) {
	query := `INSERT INTO price_vids 
	(id_1c, name_1c, registrator_id, is_active, active_change_date) 
		VALUES 
		(:id_1c, :name_1c, :registrator_id, :is_active, :active_change_date) 
		RETURNING id`

	rows, err := s.db.NamedQuery(query, size)
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

func (s *repository) Update(priceVid models.PriceVids) (int64, error) {
	if priceVid.Id == 0 {
		return 0, errors.New("id is empty")
	}

	// active_change_date и is_active не обновляем, для ручного изменения
	setQuery := `	SET id_1c = :id_1c, name_1c = :name_1c, registrator_id = :registrator_id, 
	changed_at = CURRENT_TIMESTAMP`
	query := `UPDATE price_vids` + setQuery + ` WHERE id = :id RETURNING id`

	rows, err := s.db.NamedQuery(query, priceVid)
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

func (s *repository) List() (*[]models.PriceVids, error) {
	query := `SELECT * FROM price_vids`
	var res []models.PriceVids
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *repository) GetById(id int64) (*[]models.PriceVids, error) {
	query := `SELECT * FROM price_vids WHERE id = ` + strconv.Itoa(int(id))
	var res []models.PriceVids
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
