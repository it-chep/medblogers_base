package dto

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID                 uuid.UUID      `json:"id" db:"id"`
	Name               string         `json:"name" db:"name"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at"`
	Slug               sql.NullString `json:"slug" db:"slug"`
	Body               sql.NullString `json:"body" db:"body"`
	IsActive           sql.NullBool   `json:"is_active" db:"is_active"`
	PreviewText        sql.NullString `json:"preview_text" db:"preview_text"`
	SocietyPreviewText sql.NullString `json:"society_preview" db:"society_preview"`
	AdditionalSEOText  sql.NullString `json:"additional_seo_text" db:"additional_seo_text"`
	OrderingNumber     sql.NullInt64  `json:"ordering_number" db:"ordering_number"`
}

type ValidationError struct {
	Text  string
	Field string
}

func (e ValidationError) Error() string {
	return e.Text
}
