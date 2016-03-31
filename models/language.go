package models

import (
	"encoding/json"
	"time"
)

type Language struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	IconPath   string    `json:"iconPath" db:"icon_path"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	ModifiedAt time.Time `json:"modifiedAt" db:"modified_at"`
}

func (l *Language) IsValid() error {
	return nil
}

func (l *Language) ToJson() ([]byte, error) {
	b, err := json.Marshal(l)
	return b, err
}
