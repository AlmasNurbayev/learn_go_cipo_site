package products

import (
	"cipo_cite_server/internal/models"
	"cipo_cite_server/internal/utils"
	"encoding/json"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(product models.Products) (int64, error)
	Update(product models.Products) (int64, error)
	List() (*[]models.Products, error)
}

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
	query := `SELECT products.*,
	json_agg(json_build_object('id', i.id, 'path', i.path,
	'full_name',i.full_name, 'is_main', i.is_main, 'is_active', i.is_active, 'product_id', i.product_id,
	'name', i.name)) image_registry,
	json_agg(jsonb_build_object(
	'sum', q.sum, 'qnt', q.qnt, 'store_id', q.store_id, 'size_id', q.size_id,
	'size_name_1c', q.size_name_1c)) qnt_price_registry_group,
	jsonb_build_object('id', pg.id,'name_1c', pg.name_1c) product_group,
	jsonb_build_object('id', vm.id,'name_1c', vm.name_1c) vid_modeli
	FROM products
	JOIN image_registry i ON i.product_id = products.id
	JOIN (select sum, sum(qnt) as qnt, json_agg(store_id) AS store_id, size_id, size_name_1c, product_id 
		FROM qnt_price_registry 
		GROUP BY sum, qnt, store_id, size_id, size_name_1c, product_id) q ON q.product_id = products.id 
	JOIN product_groups pg ON pg.id = products.product_group_id
	JOIN vids vm ON vm.id = products.product_group_id`
	query = utils.WhereAddParams(query, input)

	query += ` group by products.id, pg.id, vm.id`

	utils.PrintAsJSON(query)
	resSlice := []models.ProductOutput{}

	var rows, err = s.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var res models.ProductOutput
		var imageBytes []byte
		var imageStruct []models.ImageRegistryShort
		var qntPriceBytes []byte
		var qntPriceGroupStruct []models.QntPriceRegistryGroup
		var productGroupBytes []byte
		var productGroupStruct models.ProductsGroupShort
		var vidModeliBytes []byte
		var vidModeliStruct models.VidsShort

		err := rows.Scan(&res.Id, &res.Id_1c, &res.Name, &res.Name_1c, &res.Product_group_id,
			&res.Product_vid_id, &res.Registrator_id, &res.Brand_id, &res.Country_id, &res.Vid_id,
			&res.Artikul, &res.Base_ed, &res.Description, &res.Material_inside, &res.Material_podoshva,
			&res.Material_up, &res.Sex, &res.Product_folder, &res.Main_color, &res.Is_public_web,
			&res.Changed_at, &res.Created_at, &imageBytes, &qntPriceBytes,
			&productGroupBytes, &vidModeliBytes,
		)
		//err := rows.StructScan(&res)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(imageBytes, &imageStruct)
		if err != nil {
			return nil, err
		}
		res.Image_registry = imageStruct

		err = json.Unmarshal(qntPriceBytes, &qntPriceGroupStruct)
		if err != nil {
			return nil, err
		}
		res.Qnt_price_registry_group = qntPriceGroupStruct

		err = json.Unmarshal(productGroupBytes, &productGroupStruct)
		if err != nil {
			return nil, err
		}
		res.Product_group = productGroupStruct

		err = json.Unmarshal(vidModeliBytes, &vidModeliStruct)
		if err != nil {
			return nil, err
		}
		res.Vid_modeli = vidModeliStruct

		resSlice = append(resSlice, res)
	}

	utils.PrintAsJSON(resSlice)

	return &resSlice[0], nil
}
