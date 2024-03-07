package products

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

func (s *repository) Create(product models.Products) (int64, error) {
	query := `INSERT INTO products
	(id_1c, name_1c, registrator_id, name, product_group_id, product_vid_id,
		 vid_id, artikul, base_ed, description, material_inside, 
		 material_podoshva, material_up, sex, product_folder, main_color, is_public_web) 
		VALUES 
		(:id_1c, :name_1c, :registrator_id, :name, :product_group_id, :product_vid_id,
		:vid_id, :artikul, :base_ed, :description, :material_inside, 
			:material_podoshva, :material_up, :sex, :product_folder, :main_color, :is_public_web) 
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

func (s *repository) Update(product models.Products) (int64, error) {
	if product.Id == 0 {
		return 0, errors.New("id is empty")
	}
	query := `UPDATE products
	SET 
	id_1c = :id_1c, name_1c = :name_1c, registrator_id = :registrator_id, name = :name, 
	product_group_id = :product_group_id, product_vid_id = :product_vid_id,
	vid_id = :vid_id, artikul = :artikul, base_ed = :base_ed, description = :description, 
	material_inside = :material_inside, material_podoshva = :material_podoshva,
	material_up = :material_up, sex = :sex, product_folder = :product_folder,
	main_color = :main_color, is_public_web = :is_public_web, changed_at = CURRENT_TIMESTAMP
	WHERE id = :id RETURNING id`

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
	}
	return res, nil
}

func (s *repository) List() (*[]models.Products, error) {
	query := `SELECT * FROM products`
	var res []models.Products
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
