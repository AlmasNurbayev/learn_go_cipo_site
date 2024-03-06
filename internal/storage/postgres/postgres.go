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
	DatabaseUrl string
	Dbx         *sqlx.DB
	Log         *slog.Logger
}

//	func NewStore(databaseUrl string) (*PostgresStore) {
//		return &PostgresStore{
//				databaseUrl: databaseUrl,
//		}
//	}
func NewStore() *PostgresStore {
	return &PostgresStore{}
}

func (s *PostgresStore) Init(env config.Envs, log *slog.Logger) (*PostgresStore, error) {
	var databaseUrl = fmt.Sprint(
		"postgresql://",
		env.DB_USERNAME,
		":", env.DB_PASSWORD,
		"@", env.DB_HOST, ":", fmt.Sprint(env.DB_PORT), "/", env.DB_DATABASE)

	dbx, err := sqlx.Open(driverName, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "ConnectDB", err)
	}
	errPing := dbx.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("%s: %w", "PingDB", errPing)
	}
	log.Info("DB connected: " + env.DB_DATABASE)

	s.Dbx = dbx

	return s, nil
}

func (s *PostgresStore) CloseDB() error {
	return s.Dbx.Close()
}
