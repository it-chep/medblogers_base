package dto

import "github.com/google/uuid"

type Recommendation struct {
	BlogID uuid.UUID `db:"blog_id"`
	Title  string    `db:"title"`
	Slug   string    `db:"slug"`
}
