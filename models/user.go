package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

type User struct {
	Id             int64     `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	LastActivityAt time.Time `json:"lastActivityAt" db:"last_activity_at"`
}

func (u *User) IsValid() error {
	return nil
}

func UserFromForm(values url.Values) (*User, error) {
	fmt.Println(values)
	user := &User{
		Username:       values.Get("username"),
		Email:          values.Get("email"),
		Password:       values.Get("password"),
		CreatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}

	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	passwordRepeat := values.Get("password-repeat")
	if user.Password != passwordRepeat {
		return nil, errors.New("Password mismatch")
	}

	return user, nil
}

func (u *User) ToJson() ([]byte, error) {
	b, err := json.Marshal(u)
	return b, err
}
