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
type Repository struct {
	db *sqlx.Tx
}

// без транзакций
type RepositoryDb struct {
	db *sqlx.DB
}

// для транзакций
func NewRepository(db *sqlx.Tx) *Repository {
	return &Repository{
		db: db,
	}
}

// без транзакций
func NewRepositoryDb(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{
		db: db,
	}
}

func (s *Repository) Create(store models.Stores) (int64, error) {
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

func (s *Repository) Update(store models.Stores) (int64, error) {
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

func (s *Repository) List() (*[]models.Stores, error) {
	query := `SELECT * FROM stores`
	var res []models.Stores
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// без транзакций
func (s *RepositoryDb) List() (*[]models.Stores, error) {
	query := `SELECT * FROM stores WHERE is_public = true`
	var res []models.Stores
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *RepositoryDb) ListShort() (*[]models.StoresShort, error) {
	query := `SELECT id, name_1c FROM stores WHERE is_public = true`
	var res []models.StoresShort
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
