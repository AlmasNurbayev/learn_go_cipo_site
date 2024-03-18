package models

import (
	"time"
)

type Products struct {
	Id      int64  `json:"id" db:"id"`
	Id_1c   string `json:"id_1c" db:"id_1c"`
	Name_1c string `json:"name_1c" db:"name_1c"`
	Name    string `json:"name" db:"name"`

	Registrator_id   int64  `json:"registrator_id" db:"registrator_id"`
	Product_group_id int64  `json:"product_group_id" db:"product_group_id"`
	Product_vid_id   *int64 `json:"product_vid_id" db:"product_vid_id"`
	Brand_id         *int64 `json:"brand_id" db:"brand_id"`
	Country_id       *int64 `json:"country_id" db:"country_id"`
	Vid_id           *int64 `json:"vid_id" db:"vid_id"`

	Artikul           string  `json:"artikul" db:"artikul"`
	Base_ed           string  `json:"base_ed" db:"base_ed"`
	Description       *string `json:"description" db:"description"`
	Material_inside   *string `json:"material_inside" db:"material_inside"`
	Material_podoshva *string `json:"material_podoshva" db:"material_podoshva"`
	Material_up       *string `json:"material_up" db:"material_up"`
	Sex               *int    `json:"sex" db:"sex"`
	Product_folder    string  `json:"product_folder" db:"product_folder"`
	Main_color        *string `json:"main_color" db:"main_color"`
	Is_public_web     bool    `json:"is_public_web" db:"is_public_web"`

	Changed_at *time.Time `json:"changed_at" db:"changed_at"`
	Created_at time.Time  `json:"created_at" db:"created_at"`
}

type QntPriceRegistryGroup struct {
	Size_id      int64   `json:"size_id" db:"size_id"`
	Size_name_1c string  `json:"size_name_1c" db:"size_name_1c"`
	Qnt          float32 `json:"qnt" db:"qnt"`
	Sum          float32 `json:"sum" db:"sum"`
	Store_id     []int64 `json:"store_id" db:"store_id"`
}

type ProductOutput struct {
	Products
	Product_group  ProductsGroupShort `json:"product_group"`
	Vid_modeli     VidsShort          `json:"vid_modeli" db:"vid_modeli"`
	Image_registry []ImageRegistry    `json:"image_registry" db:"image_registry"`
	//Qnt_price_registry       []QntPriceRegistry      `json:"qnt_price_registry"`
	Qnt_price_registry_group []QntPriceRegistryGroup `json:"qnt_price_registry_group"`
}

type ProductNews struct {
	Id                 int64     `json:"id" db:"id"`
	Product_id         int64     `json:"product_id" db:"product_id"`
	Product_name       string    `json:"product_name" db:"product_name"`
	Sum                float32   `json:"sum" db:"sum"`
	Product_create_at  time.Time `json:"product_create_at" db:"product_create_at"`
	Name               string    `json:"name" db:"name"`
	Artikul            string    `json:"artikul" db:"artikul"`
	Description        *string   `json:"description" db:"description"`
	Material_up        *string   `json:"material_up" db:"material_up"`
	Material_inside    *string   `json:"material_inside" db:"material_inside"`
	Material_podoshva  *string   `json:"material_podoshva" db:"material_podoshva"`
	Sex                *int      `json:"sex" db:"sex"`
	Product_group_name string    `json:"product_group_name" db:"product_group_name"`
	Vid_modeli_name    string    `json:"vid_modeli_name" db:"vid_modeli_name"`
	Vid_modeli_id      int64     `json:"vid_modeli_id" db:"vid_modeli_id"`
	Image_registry     []ImageRegistryShort
	Qnt_price          []QntPriceRegistryGroup
	Image_active_path  *string `json:"image_active_path" db:"image_active_path"`
}

type ProductsList struct {
	Data          []ProductOnceForList `json:"data"`
	Full_count    int                  `json:"full_count"`
	Current_count int                  `json:"current_count"`
}

type ProductOnceForList struct {
	Sum int64 `json:"sum" db:"sum"`
	ProductNews
}
