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
	result, err := us.db.NamedQuery("INSERT INTO users (username, password, email, created_at, last_activity_at) VALUES (:username, :password, :email, :created_at, :last_activity_at) RETURNING id", user)
	if err != nil {
		return err
	}

	var id int64
	result.Scan(&id)
	user.Id = id

	return nil
}
