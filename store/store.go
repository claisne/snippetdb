package store

import "github.com/claisne/snippetdb/models"

type Store interface {
	User() UserStore
}

type UserStore interface {
	Get(id int64) (*models.User, error)
	Save(user *models.User) error
}
