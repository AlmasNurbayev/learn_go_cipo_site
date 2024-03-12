package sizes

import (
	"cipo_cite_server/internal/models"
	"errors"

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

func (s *repository) Create(size models.Sizes) (int64, error) {
	query := `INSERT INTO sizes 
	(id_1c, name_1c, registrator_id) 
		VALUES 
		(:id_1c, :name_1c, :registrator_id) 
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

func (s *repository) Update(size models.Sizes) (int64, error) {
	if size.Id == 0 {
		return 0, errors.New("id is empty")
	}
	query := `UPDATE sizes
	SET id_1c = :id_1c, name_1c = :name_1c, registrator_id = :registrator_id, changed_at = CURRENT_TIMESTAMP
	WHERE id = :id RETURNING id`

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

func (s *repository) List() (*[]models.Sizes, error) {
	query := `SELECT * FROM sizes`
	var res []models.Sizes
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// без транзакций
func (s *RepositoryDb) ListShort() (*[]models.SizesShort, error) {
	query := `SELECT id, name_1c FROM sizes`
	var res []models.SizesShort
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
