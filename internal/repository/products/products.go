package products

import (
	"cipo_cite_server/internal/models"
	"cipo_cite_server/internal/utils"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.Tx
}

func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

// без транзакций
type RepositoryDb struct {
	db *sqlx.DB
}

// без транзакций
func NewRepositoryDb(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{
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

func (s *RepositoryDb) GetByIdOrName(input map[string]interface{}, lastRegistrator int64) (*models.ProductOutput, error) {

	productsTotal := []models.ProductOutput{}

	query_products := `SELECT products.* FROM products`
	query_products = utils.WhereAddParams(query_products, input)

	err := s.db.Select(&productsTotal, query_products)
	if err != nil {
		return nil, err
	}
	if len(productsTotal) != 1 {
		return nil, errors.New("products len is not 1")
	}
	product_id := productsTotal[0].Id
	product_group_id := productsTotal[0].Product_group_id
	vid_id := productsTotal[0].Vid_id

	// делаем несколько запросов для получения остальных данных
	Image_registry := []models.ImageRegistry{}
	Image_query := `SELECT * FROM image_registry WHERE product_id = ` + strconv.Itoa(int(product_id)) + `
	ORDER BY id ASC`
	err = s.db.Select(&Image_registry, Image_query)
	if err != nil {
		return nil, err
	}

	qnt_price_registry := []models.QntPriceRegistryGroup{}
	qnt_price_query := `SELECT sum, sum(qnt) as qnt, TO_JSON(ARRAY_REMOVE(ARRAY_AGG(store_id), NULL)) as store_id, size_id, size_name_1c
	 FROM qnt_price_registry
	 WHERE 
	 product_id = ` + strconv.Itoa(int(product_id)) + ` AND 
	  registrator_id = ` + strconv.Itoa(int(lastRegistrator)) + `
		GROUP BY sum, qnt, store_id, size_id, size_name_1c
		ORDER BY size_name_1c ASC`
	rows, err := s.db.Queryx(qnt_price_query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var qnt_price_registry_item models.QntPriceRegistryGroup

		// сканируем из-за получения массива
		var storesBytes []byte
		var storesArray []int64

		err = rows.Scan(&qnt_price_registry_item.Sum, &qnt_price_registry_item.Qnt,
			&storesBytes, &qnt_price_registry_item.Size_id, &qnt_price_registry_item.Size_name_1c)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(storesBytes, &storesArray)
		if err != nil {
			return nil, err
		}
		qnt_price_registry_item.Store_id = storesArray

		qnt_price_registry = append(qnt_price_registry, qnt_price_registry_item)
	}

	product_group := []models.ProductsGroupShort{}
	product_query := `SELECT id, name_1c FROM product_groups WHERE id = ` + strconv.Itoa(int(product_group_id))
	err = s.db.Select(&product_group, product_query)
	if err != nil {
		return nil, err
	}
	if len(product_group) != 1 {
		return nil, errors.New("product_group len is not 1")
	}

	vid := []models.VidsShort{}
	vid_query := `SELECT id, name_1c FROM vids WHERE id = ` + strconv.Itoa(int(*vid_id))
	err = s.db.Select(&vid, vid_query)
	if err != nil {
		return nil, err
	}
	if len(vid) != 1 {
		return nil, errors.New("vid len is not 1")
	}

	productsTotal[0].Image_registry = Image_registry
	productsTotal[0].Product_group = product_group[0]
	productsTotal[0].Vid_modeli = vid[0]
	productsTotal[0].Qnt_price_registry_group = qnt_price_registry

	return &productsTotal[0], nil
}
