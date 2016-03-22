package store

import "github.com/claisne/snippetdb/models"

type SqlUserStore struct {
	*SqlStore
}

func NewSqlUserStore(sqlStore *SqlStore) UserStore {
	us := &SqlUserStore{sqlStore}
	return us
}

func (us *SqlUserStore) Get(id int64) (*models.User, error) {
	user := &models.User{}
	err := us.db.Get(user, "SELECT * FROM users WHERE id=$1", id)
	return user, err
}

func (us *SqlUserStore) Save(user *models.User) error {
	return nil
}
