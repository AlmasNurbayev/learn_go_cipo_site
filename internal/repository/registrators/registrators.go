package registrators

import (
	"cipo_cite_server/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(registrator models.RegistratorsModel) (int64, error)
	List() (*[]models.RegistratorsModel, error)
}

type repository struct {
	db *sqlx.Tx
}

func NewRepository(db *sqlx.Tx) *repository {
	return &repository{
		db: db,
	}
}

func (s *repository) Create(registrator models.RegistratorsModel) (int64, error) {
	query := `INSERT INTO registrators
	(operation_date,name_folder,name_file,user_id,date_schema,
		id_catalog,id_class,name_catalog,name_class,ver_schema,is_only_change,changed_at) 
		VALUES 
		(:operation_date,:name_folder,:name_file,:user_id,:date_schema,:id_catalog,
			:id_class,:name_catalog,:name_class,:ver_schema,:is_only_change,:changed_at) 
		RETURNING *`

	rows, err := s.db.NamedQuery(query, registrator)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	res := models.RegistratorsModel{}

	for rows.Next() {
		err := rows.StructScan(&res)
		if err != nil {
			return 0, err
		}
	}
	return res.Id, nil
}

func (s *repository) List() (*[]models.RegistratorsModel, error) {
	query := `SELECT * FROM registrators`
	var res []models.RegistratorsModel
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *repository) GetById(id int64) (*[]models.RegistratorsModel, error) {
	query := `SELECT * FROM registrators WHERE id = ` + strconv.Itoa(int(id))
	var res []models.RegistratorsModel
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
