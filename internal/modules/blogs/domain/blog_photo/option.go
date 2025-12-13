package blog_photo

import "github.com/google/uuid"

// Option определяет тип функции для установки опций BlogPhoto.
type Option func(bp *BlogPhoto)

// WithPhotoID устанавливает ID фотографии.
func WithPhotoID(id uuid.UUID) Option {
	return func(bp *BlogPhoto) {
		bp.id = id
	}
}

// WithBlogID устанавливает ID блога, к которому относится фотография.
func WithBlogID(blogID uuid.UUID) Option {
	return func(bp *BlogPhoto) {
		bp.blogID = blogID
	}
}

// WithIsPrimary устанавливает, является ли фотография основной.
func WithIsPrimary(isPrimary bool) Option {
	return func(bp *BlogPhoto) {
		bp.isPrimary = isPrimary
	}
}

// WithFileType устанавливает тип файла фотографии.
func WithFileType(fileType string) Option {
	return func(bp *BlogPhoto) {
		bp.fileType = fileType
	}
}
