package blog_photo

import "github.com/google/uuid"

// GetID возвращает ID фотографии.
func (bp *BlogPhoto) GetID() uuid.UUID {
	return bp.id
}

// GetBlogID возвращает ID блога.
func (bp *BlogPhoto) GetBlogID() uuid.UUID {
	return bp.blogID
}

// GetIsPrimary возвращает, является ли фотография основной.
func (bp *BlogPhoto) GetIsPrimary() bool {
	return bp.isPrimary
}

// IsPrimary - альтернативное название для метода GetIsPrimary.
func (bp *BlogPhoto) IsPrimary() bool {
	return bp.isPrimary
}

// GetFileType возвращает тип файла фотографии.
func (bp *BlogPhoto) GetFileType() string {
	return bp.fileType
}
