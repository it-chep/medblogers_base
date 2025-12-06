package blog

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	id                 uuid.UUID
	title              string
	slug               string
	body               string
	previewText        string
	societyPreviewText string
	additionalSEOText  string
	orderingNumber     int64
	createdAt          time.Time

	isActive        bool
	primaryPhotoURL string
}

type Blogs []*Blog

func New(options ...Option) *Blog {
	d := &Blog{}
	for _, option := range options {
		option(d)
	}
	return d
}
