package offer

import (
	"time"

	"github.com/google/uuid"
)

type Offer struct {
	id                   uuid.UUID
	cooperationTypeID    int64
	businessCategoryID   int64
	title                string
	description          string
	price                int64
	contentFormatID      int64
	brandID              int64
	publicationDate      *time.Time
	adMarkingResponsible string
	responsesCapacity    int64
	isActive             bool
	createdAt            time.Time
}

func New(options ...Option) *Offer {
	item := &Offer{}
	for _, option := range options {
		option(item)
	}

	return item
}
