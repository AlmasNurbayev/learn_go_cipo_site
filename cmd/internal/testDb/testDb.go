package testdb

import (
	"cipo_cite_server/cmd/internal/storage/store"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func Testdb(dbx *sqlx.DB, log *slog.Logger) {
	res, err := store.CreateStore(store.Model{
		Name_1c: time.Now().String(),
		Id_1c:   uuid.New().String()},
		dbx)
	if err != nil {
		log.Error("cannot create store: " + err.Error())
	}
	log.Info("store created: ", res)

	res2, err := store.GetStore(dbx)
	if err != nil {
		log.Error("cannot get store: " + err.Error())
	}
	log.Info("stores get: ", res2)
}
