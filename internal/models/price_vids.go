package models

import (
	"time"
)

type PriceVids struct {
	Id      int64  `json:"id" db:"id"`
	Name_1c string `json:"name_1c" db:"name_1c"`
	Id_1c   string `json:"id_1c" db:"id_1c"`

	Is_active          bool      `json:"is_active" db:"is_active"`
	Active_change_date time.Time `json:"active_change_date" db:"active_change_date"`

	Registrator_id int64 `json:"registrator_id" db:"registrator_id"`

	Changed_at *time.Time `json:"changed_at" db:"changed_at"`
	Created_at time.Time  `json:"created_at" db:"created_at"`
}
