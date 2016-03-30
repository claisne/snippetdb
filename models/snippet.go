package models

import (
	"encoding/json"
	"net/url"
	"time"
)

type Snippet struct {
	Id             int64     `json:"id"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	LastActivityAt time.Time `json:"lastActivityAt" db:"last_activity_at"`
}

func (s *Snippet) IsValid() error {
	return nil
}

func NewSnippetFromForm(values url.Values) (*Snippet, error) {
	snippet := &Snippet{
		CreatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}

	return snippet
}

func (u *Snippet) ToJson() ([]byte, error) {
	b, err := json.Marshal(u)
	return b, err
}
