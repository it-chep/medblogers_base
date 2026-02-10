package dto

import "github.com/google/uuid"

type ImageDTO struct {
	ImageID   uuid.UUID `json:"id" db:"id"`
	BlogID    uuid.UUID `json:"blog_id" db:"blog_id"`
	FileType  string    `json:"file_type" db:"file_type"`
	IsPrimary bool      `json:"is_primary" db:"is_primary"`
}
