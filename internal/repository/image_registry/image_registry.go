package image_registry

import (
	"cipo_cite_server/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(product models.Products) (int64, error)
	Update(product models.Products) (int64, error)
	List() (*[]models.Products, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(product models.ImageRegistry) (int64, error) {
	query := `INSERT INTO image_registry
	(is_main, resolution, size, full_name, name, path, operation_date, is_active, registrator_id, 
		product_id, changed_at ) 
		VALUES 
		(:is_main, :resolution, :size, :full_name, :name, :path, :operation_date, :is_active, :registrator_id, 
			:product_id, :changed_at) 
		RETURNING id`

	rows, err := s.db.NamedQuery(query, product)
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

func (s *repository) Update(image models.ImageRegistry) (int64, error) {
	if image.Name == "" {
		return 0, errors.New("Unique name is empty")
	}
	query := `UPDATE image_registry
	SET 
	is_main =:is_main, resolution = :resolution, size = :size, full_name = :full_name, path = :path,
	operation_date = :operation_date, is_active = :is_active, registrator_id = :registrator_id, 
	product_id =:product_id, changed_at = CURRENT_TIMESTAMP
	WHERE name = :name RETURNING id`

	rows, err := s.db.NamedQuery(query, image)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var res int64

	for rows.Next() {
		err := rows.Scan(&res)
		if err != nil {
			return res, err
		}
	}
	return 0, nil
}

func (s *repository) List() (*[]models.ImageRegistry, error) {
	query := `SELECT * FROM image_registry`
	var res []models.ImageRegistry
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
