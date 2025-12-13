package blog_photo

import "github.com/google/uuid"

type BlogPhoto struct {
	id        uuid.UUID
	blogID    uuid.UUID
	isPrimary bool
	fileType  string
}

func New(options ...Option) *BlogPhoto {
	d := &BlogPhoto{}
	for _, option := range options {
		option(d)
	}
	return d
}
