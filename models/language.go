package models

import (
	"encoding/json"
	"net/url"
	"time"
)

type Language struct {
	Id             int64     `json:"id"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	LastActivityAt time.Time `json:"lastActivityAt" db:"last_activity_at"`
}

func (l *Language) IsValid() error {
	return nil
}

func NewLanguageFromForm(values url.Values) (*Language, error) {
	language := &Language{
		CreatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}

	return language
}

func (l *Language) ToJson() ([]byte, error) {
	b, err := json.Marshal(u)
	return b, err
}
