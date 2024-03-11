package news

import (
	"cipo_cite_server/internal/models"
	"cipo_cite_server/internal/server/handlers/news/newsFilter"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// без транзакций
type RepositoryDb struct {
	db *sqlx.DB
}

// без транзакций
func NewRepositoryDb(db *sqlx.DB) *RepositoryDb {
	return &RepositoryDb{
		db: db,
	}
}

// без транзакции
func (s *RepositoryDb) List(filter *newsFilter.Filter) (*[]models.News, error) {
	//utils.PrintAsJSON(filter)
	query := "SELECT * FROM news ORDER BY id DESC"
	if filter.News != "" {
		query = query + " LIMIT " + filter.News + ";"
	} else {
		query = query + " LIMIT 5"
	}
	fmt.Println(query)

	var res []models.News
	var err = s.db.Select(&res, query)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
