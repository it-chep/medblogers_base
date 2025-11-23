package dto

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Request struct {
	Name               string
	Slug               string
	Body               string
	IsActive           bool
	PreviewText        string
	SocietyPreviewText string
	AdditionalSEOText  string
	OrderingNumber     int64
}

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
