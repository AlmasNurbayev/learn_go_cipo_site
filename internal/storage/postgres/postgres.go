package postgres

import (
	"fmt"

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
			return nil, fmt.Errorf("%s: %w", "ConnectDB" ,err)
	}
	errPing := dbx.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("%s: %w", "PingDB" ,errPing)
	}	

	s.dbx = dbx
	return dbx, nil
}

func (s *PostgresStore) CloseDB() error {
	return s.dbx.Close()
}