package storage

import (
	"embed"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func MigrationsUp(db *sqlx.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		db.Close()
		panic(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		db.Close()
		panic(err)
	}
}
