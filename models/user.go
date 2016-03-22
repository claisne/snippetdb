package models

import "encoding/json"

type User struct {
	Id             int64  `json:"id"`
	CreatedAt      int64  `json:"createdAt"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	LastActivityAt int64  `json:"lastActivityAt"`
}

func (u *User) IsValid() bool {
	return false
}

func (u *User) ToJson() ([]byte, error) {
	b, err := json.Marshal(u)
	return b, err
}
