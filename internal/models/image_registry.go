package models

import (
	"time"
)

type ImageRegistry struct {
	Id         int64   `json:"id" db:"id"`
	Is_main    bool    `json:"is_main" db:"is_main"`
	Resolution *string `json:"resolution" db:"resolution"`
	Size       int     `json:"size" db:"size"`
	Full_name  string  `json:"full_name" db:"full_name"`
	Name       string  `json:"name" db:"name"`
	Path       string  `json:"path" db:"path"`

	Operation_date time.Time `json:"operation_date" db:"operation_date"`
	Is_active      bool      `json:"is_active" db:"is_active"`

	Registrator_id int64 `json:"registrator_id" db:"registrator_id"`
	Product_id     int64 `json:"product_id" db:"product_id"`

	Changed_at *time.Time `json:"changed_at" db:"changed_at"`
	Created_at time.Time  `json:"created_at" db:"created_at"`
}

type ImageRegistryShort struct {
	Id         int64  `json:"id" db:"id"`
	Is_main    bool   `json:"is_main" db:"is_main"`
	Full_name  string `json:"full_name" db:"full_name"`
	Name       string `json:"name" db:"name"`
	Path       string `json:"path" db:"path"`
	Is_active  bool   `json:"is_active" db:"is_active"`
	Product_id int64  `json:"product_id" db:"product_id"`
}
