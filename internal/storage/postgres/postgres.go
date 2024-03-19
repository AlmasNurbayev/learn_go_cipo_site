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
	Tx          *sqlx.Tx
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

	dbx, err := sqlx.Connect(driverName, databaseUrl)
	dbx.SetMaxOpenConns(10) // Максимальное количество открытых соединений
	dbx.SetMaxIdleConns(5)  // Максимальное количество простаивающих соединений
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "ConnectDB", err)
	}
	errPing := dbx.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("%s: %w", "PingDB", errPing)
	}
	tx, err := dbx.Beginx()
	if err != nil {
		return nil, err
	}
	s.Dbx = dbx
	s.Tx = tx
	log.Info("DB connected: " + env.DB_DATABASE)

	return s, nil
}

func (s *PostgresStore) CloseDB() error {
	return s.Dbx.Close()
}
