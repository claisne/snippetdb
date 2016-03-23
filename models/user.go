package models

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                  int64     `json:"id"`
	Username            string    `json:"username"`
	Password            string    `json:"password"`
	Email               string    `json:"email"`
	SnippetViewsCount   int64     `json:"snippetViewsCount" db:"snippet_views_count"`
	SnippetUpvotesCount int64     `json:"snippetUpvotesCount" db:"snippet_upvotes_count"`
	SnippetSavesCount   int64     `json:"snippetSavesCount" db:"snippet_saves_count"`
	CreatedAt           time.Time `json:"createdAt" db:"created_at"`
	LastActivityAt      time.Time `json:"lastActivityAt" db:"last_activity_at"`
}

func init() {
	gob.Register(&User{})
}

func (u *User) IsValid() error {
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func NewUserFromForm(values url.Values) (*User, error) {
	user := &User{
		Username:            values.Get("username"),
		Password:            values.Get("password"),
		Email:               values.Get("email"),
		CreatedAt:           time.Now(),
		SnippetViewsCount:   0,
		SnippetSavesCount:   0,
		SnippetUpvotesCount: 0,
		LastActivityAt:      time.Now(),
	}

	// Check if valid attributes
	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	// Check password confirmation
	passwordRepeat := values.Get("password-repeat")
	if user.Password != passwordRepeat {
		return nil, errors.New("Password mismatch")
	}

	// Hash password
	hashedPassword, err := hashPassword(user.Password)
	if user.Password != passwordRepeat {
		return nil, errors.New("Failed to hash")
	}

	user.Password = hashedPassword

	return user, nil
}

func (u *User) ToJson() ([]byte, error) {
	b, err := json.Marshal(u)
	return b, err
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
