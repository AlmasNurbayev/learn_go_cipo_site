package store

import (
	"cipo_cite_server/internal/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// type StorePostgres struct {
// 	db *sqlx.DB
// }

// func NewStorePostgres(db *sqlx.DB) *StorePostgres {
// 	return &StorePostgres{db: db}
// }

func CreateStore(store Model, db *sqlx.DB) (Model, error) {

	res := Model{}

	query := "INSERT INTO store (id_1c, name_1c) values (:id_1c, :name_1c) RETURNING *"
	rows, err := db.NamedQuery(query, store)

	for rows.Next() {
		if err := rows.StructScan(&res); err != nil {
			return res, err
		}
	}
	return res, err
}

func GetStore(db *sqlx.DB) ([]Model, error) {

	res := []Model{}

	query := "SELECT * FROM STORE"
	//
	params := map[string]interface{}{
		"id":      []int{1, 7},
		"name_1c": "name",
		"id_1c":   nil,
	}

	fmt.Println("params", utils.WhereAddParams(query, params))
	rows, err := db.NamedQuery(utils.WhereAddParams(query, params), params)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		resRow := Model{}
		if err := rows.StructScan(&resRow); err != nil {
			return res, err
		}

		res = append(res, resRow)
	}
	return res, err
}
