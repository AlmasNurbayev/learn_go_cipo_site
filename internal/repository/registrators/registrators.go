package registrators

import (
	"cipo_cite_server/internal/models"

	"github.com/jmoiron/sqlx"
)

type Operations interface {
	Create(registrator models.RegistratorsModel) (int64, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
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
