package registrators

import (
	"cipo_cite_server/internal/models"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func Insert(db *sqlx.DB, log *slog.Logger, tableName string, data models.RegistratorsModel) (interface{}, error) {
	return operations.Insert(db, log, tableName, data)
}
