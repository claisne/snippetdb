package models

import (
	"encoding/json"
	"net/url"
	"time"
)

type Snippet struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"userId" db:"user_id"`
	LanguageId   int64     `json:"languageId" db:"language_id"`
	Title        string    `json:"title"`
	Code         string    `json:"code"`
	ViewsCount   int64     `json:"viewsCount" db:"views_count"`
	UpvotesCount int64     `json:"upvotesCount" db:"upvotes_count"`
	SavesCount   int64     `json:"savesCount" db:"saves_count"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	ModifiedAt   time.Time `json:"modifiedAt" db:"modified_at"`
}

func (s *Snippet) IsValid() error {
	return nil
}

func NewSnippetFromForm(values url.Values) (*Snippet, error) {
	snippet := &Snippet{
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	return snippet, nil
}

func (u *Snippet) ToJson() ([]byte, error) {
	b, err := json.Marshal(u)
	return b, err
}
