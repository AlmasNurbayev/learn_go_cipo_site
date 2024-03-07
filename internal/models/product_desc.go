package models

import (
	"time"
)

type ProductsDesc struct {
	Id             int64      `json:"id" db:"id"`
	Id_1c          string     `json:"id_1c" db:"id_1c"`
	Name_1c        string     `json:"name_1c" db:"name_1c"`
	Field          string     `json:"field" db:"field"`
	Registrator_id int64      `json:"registrator_id" db:"registrator_id"`
	Changed_at     *time.Time `json:"changed_at" db:"changed_at"`
	Created_at     time.Time  `json:"created_at" db:"created_at"`
}
