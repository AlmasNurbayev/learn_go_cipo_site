package products

import (
	"cipo_cite_server/internal/models"
	"cipo_cite_server/internal/server/handlers/product/productFilter"
	"cipo_cite_server/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

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

func (s *RepositoryDb) GetProductNews(count int, lastRegistrator int64) (*[]models.ProductNews, error) {
	productsNews := []models.ProductNews{}
	query := `
	SELECT pg.name_1c as product_group_name, v.name_1c as vid_modeli_name, p.id, p.name, p.artikul, p.description, 
p.material_podoshva, p.material_up, p.material_inside, p. sex, p.created_at as product_create_at ,
to_json(array_agg(q2)) as qnt_price,
to_json(images.agg) as image_registry
FROM qnt_price_registry q
join lateral (select array_agg( store_id) as store_id, size_id, sum, qnt, size_name_1c from qnt_price_registry 
where id = q.id group by store_id, sum, qnt, id, size_id ) as q2 on true 
join products p on p.id = product_id
join product_groups pg on p.product_group_id = pg.id
join vids v on p.vid_id = v.id
join lateral (select (array_agg( jsonb_build_object('name', image_registry.name, 'full_name', full_name, 'is_active', is_active, 'is_main', is_main))) 
as agg from image_registry where image_registry.product_id = p.id) as images on true
where q.registrator_id = ` + strconv.Itoa(int(lastRegistrator)) + ` and q.qnt > 0
GROUP by p.name, p.id, images.agg, pg.name_1c, v.name_1c
order by product_create_at desc
limit ` + strconv.Itoa(count)

	rows, err := s.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var productNews_item models.ProductNews

		// сканируем из-за получения массива
		var imagesBytes []byte
		var imagesArray []models.ImageRegistryShort

		var qntBytes []byte
		var qntArray []models.QntPriceRegistryGroup

		err = rows.Scan(&productNews_item.Product_group_name, &productNews_item.Vid_modeli_name,
			&productNews_item.Id, &productNews_item.Name, &productNews_item.Artikul, &productNews_item.Description,
			&productNews_item.Material_podoshva, &productNews_item.Material_up, &productNews_item.Material_inside,
			&productNews_item.Sex, &productNews_item.Product_create_at, &qntBytes, &imagesBytes)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(imagesBytes, &imagesArray)
		if err != nil {
			return nil, err
		}
		productNews_item.Image_registry = imagesArray

		err = json.Unmarshal(qntBytes, &qntArray)
		if err != nil {
			return nil, err
		}
		productNews_item.Qnt_price = qntArray

		if len(productNews_item.Image_registry) > 0 {
			indexMain := slices.IndexFunc(productNews_item.Image_registry, func(i models.ImageRegistryShort) bool {
				return i.Is_main
			})
			if indexMain != -1 {
				productNews_item.Image_active_path = &productNews_item.Image_registry[indexMain].Full_name
			}
		}

		productsNews = append(productsNews, productNews_item)

	}
	return &productsNews, nil
}

func (s *RepositoryDb) List(filters productFilter.FilterListT, lastRegistrator int64) (*models.ProductsList, error) {
	productList := []models.ProductOnceForList{}

	var orderString, whereString string
	for key, value := range filters.Base.Sort {
		orderString = key + " " + value
	}
	if orderString != "" {
		orderString = "ORDER BY " + orderString
	}

	fmt.Println(filters.Filters["Size"])
	if filters.Filters["size"] != nil {
		whereString = whereString + " AND q2.size_id = ANY (:size) "
	}
	if filters.Filters["product_group"] != nil {
		whereString = whereString + " AND pg.id = ANY (:product_group) "
	}
	if filters.Filters["search_name"] != nil {
		whereString = whereString + " AND p.name ILIKE :search_name "
	}

	query := `SELECT q2.sum, pg.name_1c as product_group_name, v.name_1c as vid_modeli_name, p.id, p.name, p.artikul, p.description, 
	p.material_podoshva, p.material_up, p.material_inside, p. sex, p.created_at as product_create_at ,
	to_json(array_agg(q2)) as qnt_price,
	to_json(images.agg) as image_registry
	FROM qnt_price_registry q
	join lateral (select array_agg( store_id) as store_id, size_id, sum, qnt, size_name_1c from qnt_price_registry 
	where id = q.id group by store_id, sum, qnt, id, size_id ) as q2 on true 
	join products p on p.id = product_id
	join product_groups pg on p.product_group_id = pg.id
	join vids v on p.vid_id = v.id
	join lateral (select (array_agg( jsonb_build_object('name', image_registry.name, 'full_name', full_name, 'is_active', is_active, 'is_main', is_main))) 
	as agg from image_registry where image_registry.product_id = p.id) as images on true
	WHERE q.registrator_id = ` + strconv.Itoa(int(lastRegistrator)) + whereString + ` AND q.qnt > 0
	GROUP by p.name, p.id, images.agg, pg.name_1c, v.name_1c, q2.sum
	 ` + orderString
	queryWithFilters := query +
		` limit ` + strconv.Itoa(filters.Base.Take) + ` offset ` + strconv.Itoa(filters.Base.Skip)

	fmt.Println(queryWithFilters)
	utils.PrintAsJSON(filters)

	rows, err := s.db.NamedQuery(queryWithFilters, filters.Filters)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var productNews_item models.ProductOnceForList

		// сканируем из-за конвертации байтов в массивы
		var imagesBytes []byte
		var imagesArray []models.ImageRegistryShort

		var qntBytes []byte
		var qntArray []models.QntPriceRegistryGroup

		var sumString string

		err = rows.Scan(&sumString, &productNews_item.Product_group_name, &productNews_item.Vid_modeli_name,
			&productNews_item.Id, &productNews_item.Name, &productNews_item.Artikul, &productNews_item.Description,
			&productNews_item.Material_podoshva, &productNews_item.Material_up, &productNews_item.Material_inside,
			&productNews_item.Sex, &productNews_item.Product_create_at, &qntBytes, &imagesBytes)
		if err != nil {
			return nil, err
		}

		f, err := strconv.ParseFloat(strings.TrimSpace(sumString), 64)
		if err != nil {
			return nil, err
		}
		productNews_item.Sum = int64(math.Round(f))

		err = json.Unmarshal(imagesBytes, &imagesArray)
		if err != nil {
			return nil, err
		}
		productNews_item.Image_registry = imagesArray

		err = json.Unmarshal(qntBytes, &qntArray)
		if err != nil {
			return nil, err
		}
		productNews_item.Qnt_price = qntArray

		if len(productNews_item.Image_registry) > 0 {
			indexMain := slices.IndexFunc(productNews_item.Image_registry, func(i models.ImageRegistryShort) bool {
				return i.Is_main
			})
			if indexMain != -1 {
				productNews_item.Image_active_path = &productNews_item.Image_registry[indexMain].Full_name
			}
		}

		productList = append(productList, productNews_item)

	}
	utils.PrintAsJSON(filters)

	return &models.ProductsList{
		Data:          productList,
		Full_count:    0,
		Current_count: len(productList)}, nil
}
