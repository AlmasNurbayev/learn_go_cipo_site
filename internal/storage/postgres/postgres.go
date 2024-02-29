package postgres

import (
	"fmt"
	"log/slog"

	"cipo_cite_server/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const driverName = "pgx"

type PostgresStore struct {
	databaseUrl string
	dbx         *sqlx.DB
}

// func NewStore(databaseUrl string) (*PostgresStore) {
// 	return &PostgresStore{
// 			databaseUrl: databaseUrl,
// 	}
// }

func (s *PostgresStore) New(databaseUrl string) (*sqlx.DB, error) {
	dbx, err := sqlx.Open(driverName, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "ConnectDB", err)
	}
	errPing := dbx.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("%s: %w", "PingDB", errPing)
	}

	s.dbx = dbx
	s.databaseUrl = databaseUrl
	return dbx, nil
}

func (s *PostgresStore) CloseDB() error {
	return s.dbx.Close()
}

func InitPostgresStore(env config.Envs, log *slog.Logger) *sqlx.DB {
	postgresStore := PostgresStore{}
	var databaseUrl = fmt.Sprint(
		"postgresql://",
		env.DB_USERNAME,
		":", env.DB_PASSWORD,
		"@", env.DB_HOST, ":", fmt.Sprint(env.DB_PORT), "/", env.DB_DATABASE)
	dbx, err := postgresStore.New(databaseUrl)
	dbx.DB.Ping()
	if err != nil {
		log.Error("cannot connect to database: " + err.Error())
		panic(err)
	}
	log.Info("DB connected: " + env.DB_DATABASE)
	return dbx
}
