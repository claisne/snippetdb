package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SqlStore struct {
	db   *sqlx.DB
	user UserStore
}

func NewSqlStore(driver string, dsn string) (Store, error) {
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		return nil, err
	}

	sqlStore := &SqlStore{db: db}
	sqlStore.user = NewSqlUserStore(sqlStore)

	return sqlStore, nil
}

func (ss SqlStore) User() UserStore {
	return ss.user
}
