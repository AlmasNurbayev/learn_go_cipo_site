package qnt_price_registry

import (
	"cipo_cite_server/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(product models.QntPriceRegistry) (int64, error)
	List() (*[]models.QntPriceRegistry, error)
	GetByRegistratorId(int64) (*[]models.QntPriceRegistry, error)
}

type repository struct {
	db *sqlx.Tx
}

func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(qntPrice models.QntPriceRegistry) (int64, error) {

	query := `INSERT INTO qnt_price_registry
	(sum, qnt, operation_date, discount_percent, discount_begin, discount_end, store_id,
		product_id, price_vid_id, size_id, registrator_id, product_group_id, vid_modeli_id,
		size_name_1c, product_name, product_created_at, changed_at)
		VALUES 
		(:sum, :qnt, :operation_date, :discount_percent, :discount_begin, :discount_end, :store_id,
			:product_id, :price_vid_id, :size_id, :registrator_id, :product_group_id, :vid_modeli_id,
			:size_name_1c, :product_name, :product_created_at, :changed_at) 
		RETURNING id`

	rows, err := s.db.NamedQuery(query, qntPrice)
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

func (s *repository) List() (*[]models.QntPriceRegistry, error) {
	query := `SELECT * FROM qnt_price_registry`
	var res []models.QntPriceRegistry
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *repository) GetByRegistratorId(id int64) (*[]models.QntPriceRegistry, error) {
	query := `SELECT * FROM qnt_price_registry WHERE registrator_id = ` + strconv.Itoa(int(id))
	var res []models.QntPriceRegistry
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
