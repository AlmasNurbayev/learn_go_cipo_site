package stores

import (
	"cipo_cite_server/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(store models.Stores) (int64, error)
	Update(store models.Stores) (int64, error)
	List() (*[]models.Stores, error)
}

// для транзакций
type repository struct {
	db *sqlx.Tx
}

// без транзакций
type repositoryDb struct {
	db *sqlx.DB
}

// для транзакций
func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

// без транзакций
func NewRepositoryDb(db *sqlx.DB) *repositoryDb {
	return &repositoryDb{
		db: db,
	}
}

func (s *repository) Create(store models.Stores) (int64, error) {
	// записываем только часть полей - остальные для правки вручную
	query := `INSERT INTO stores 
	(id_1c, name_1c, registrator_id) 
		VALUES 
		(:id_1c, :name_1c, :registrator_id) 
		RETURNING id`

	rows, err := s.db.NamedQuery(query, store)
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

func (s *repository) Update(store models.Stores) (int64, error) {
	if store.Id == 0 {
		return 0, errors.New("id is empty")
	}

	// записываем только часть полей - остальные для правки вручную
	query := `UPDATE stores
	SET id_1c = :id_1c, name_1c = :name_1c, registrator_id = :registrator_id, changed_at = CURRENT_TIMESTAMP
	WHERE id = :id RETURNING id`

	rows, err := s.db.NamedQuery(query, store)
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

// для транзакции
func (s *repository) List() (*[]models.Stores, error) {
	query := `SELECT * FROM stores`
	var res []models.Stores
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// без транзакции
func (s *repositoryDb) List() (*[]models.Stores, error) {
	query := `SELECT * FROM stores`
	var res []models.Stores
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
