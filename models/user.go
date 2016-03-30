package models

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	valid "github.com/asaskevich/govalidator"
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
	if !valid.IsAlphanumeric(u.Username) {
		return errors.New("The username must only contains alphanumerical characters")
	}

	if len(u.Username) < 3 || len(u.Username) > 10 {
		return errors.New("The username must have a maximun of 10 characters and a minimum of 3")
	}

	if u.Email != "" && !valid.IsEmail(u.Email) {
		return errors.New("The email must be empty or a valid adress")
	}

	if len(u.Email) > 150 {
		return errors.New("The email must have a maximum of 150 characters")
	}

	if len(u.Password) < 4 || len(u.Password) > 50 {
		return errors.New("The username must have a maximun of 50 characters and a minimum of 4")
	}

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
		return nil, errors.New("Internal Error. Please retry")
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
