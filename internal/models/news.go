package models

import (
	"time"
)

type News struct {
	Id         int64      `json:"id" db:"id"`
	Title      string     `json:"title" db:"title"`
	Data       string     `json:"data" db:"data"`
	Image_path *string    `json:"image_path" db:"image_path"`
	Changed_at *time.Time `json:"changed_at" db:"changed_at"`
	Created_at time.Time  `json:"created_at" db:"created_at"`
}
