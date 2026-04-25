package dto

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Blog struct {
	BlogID         uuid.UUID     `json:"id" db:"id"`
	Name           string        `json:"name" db:"name"`
	IsActive       sql.NullBool  `json:"is_active" db:"is_active"`
	OrderingNumber sql.NullInt64 `json:"ordering_number" db:"ordering_number"`
	ViewsCount     int64         `json:"views_count"`
}

type Blogs []Blog

func (b Blogs) GetIDs() []uuid.UUID {
	return lo.Map(b, func(item Blog, _ int) uuid.UUID {
		return item.BlogID
	})
}
