package models

import (
	"time"
)

type AutoId struct {
	ID uint `json:"id" db:"id"`
}

type ModelTimestamps struct {
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
