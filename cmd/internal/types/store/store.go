package store

type Model struct {
	Id       int    `json:"id" db:"id"`
	Id_1c    string `json:"id_1c" binding:"required"`
	Name_1c	 string `json:"name_1c" binding:"required"`
}