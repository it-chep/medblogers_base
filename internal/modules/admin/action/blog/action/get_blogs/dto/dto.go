package dto

import (
	"database/sql"
	"github.com/google/uuid"
)

type Blog struct {
	BlogID   uuid.UUID    `json:"id" db:"id"`
	Name     string       `json:"name" db:"name"`
	IsActive sql.NullBool `json:"is_active" db:"is_active"`
}

type Blogs []Blog
