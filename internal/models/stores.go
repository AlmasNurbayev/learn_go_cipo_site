package models

import (
	"time"
)

type Stores struct {
	Id             int64  `json:"id" db:"id"`
	Id_1c          string `json:"id_1c" db:"id_1c"`
	Name_1c        string `json:"name_1c" db:"name_1c"`
	Registrator_id int64  `json:"registrator_id" db:"registrator_id"`

	Address    *string `json:"address" db:"address"`
	Link_2gis  *string `json:"link_2gis" db:"link_2gis"`
	Phone      *string `json:"phone" db:"phone"`
	City       *string `json:"city" db:"city"`
	Image_path *string `json:"image_path" db:"image_path"`
	Is_public  *string `json:"is_public" db:"is_public"`

	Changed_at *time.Time `json:"changed_at" db:"changed_at"`
	Created_at time.Time  `json:"created_at" db:"created_at"`
}
