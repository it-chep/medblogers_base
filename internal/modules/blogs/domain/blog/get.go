package blog

import (
	"time"

	"github.com/samber/lo"

	"github.com/google/uuid"
)

// GetID возвращает ID блога.
func (b *Blog) GetID() uuid.UUID {
	return b.id
}

// GetTitle возвращает заголовок блога.
func (b *Blog) GetTitle() string {
	return b.title
}

// GetSlug возвращает slug блога.
func (b *Blog) GetSlug() string {
	return b.slug
}

// GetBody возвращает тело блога.
func (b *Blog) GetBody() string {
	return b.body
}

// GetPreviewText возвращает превью-текст блога.
func (b *Blog) GetPreviewText() string {
	return b.previewText
}

// GetSocietyPreviewText возвращает превью-текст для соц сетей
func (b *Blog) GetSocietyPreviewText() string {
	return b.societyPreviewText
}

// GetAdditionalSEOText возвращает дополнительный SEO-текст.
func (b *Blog) GetAdditionalSEOText() string {
	return b.additionalSEOText
}

// GetOrderingNumber возвращает номер для сортировки.
func (b *Blog) GetOrderingNumber() int64 {
	return b.orderingNumber
}

// GetCreatedAt возвращает время создания.
func (b *Blog) GetCreatedAt() time.Time {
	return b.createdAt
}

// GetIsActive возвращает статус активности.
func (b *Blog) GetIsActive() bool {
	return b.isActive
}

// IsActive .
func (b *Blog) IsActive() bool {
	return b.isActive
}

// GetDoctorID .
func (b *Blog) GetDoctorID() int64 {
	return b.doctorID
}

// HasAuthor .
func (b *Blog) HasAuthor() bool {
	return b.doctorID > 0
}

// GetPrimaryPhotoURL возвращает главную фотку
func (b *Blog) GetPrimaryPhotoURL() string {
	return b.primaryPhotoURL
}

/////////////////////////////////////////////
///////////////////BLOGS/////////////////////
/////////////////////////////////////////////

func (b Blogs) GetIDs() []uuid.UUID {
	return lo.Map(b, func(blog *Blog, i int) uuid.UUID {
		return blog.id
	})
}
