package blog

import (
	"time"

	"github.com/google/uuid"
)

// Option определяет тип функции для установки опций Blog.
type Option func(b *Blog)

// WithID устанавливает ID блога.
func WithID(id uuid.UUID) Option {
	return func(b *Blog) {
		b.id = id
	}
}

// WithTitle устанавливает заголовок блога.
func WithTitle(title string) Option {
	return func(b *Blog) {
		b.title = title
	}
}

// WithSlug устанавливает slug блога.
func WithSlug(slug string) Option {
	return func(b *Blog) {
		b.slug = slug
	}
}

// WithBody устанавливает тело блога.
func WithBody(body string) Option {
	return func(b *Blog) {
		b.body = body
	}
}

// WithPreviewText устанавливает превью-текст блога.
func WithPreviewText(previewText string) Option {
	return func(b *Blog) {
		b.previewText = previewText
	}
}

// WithSocietyPreviewText устанавливает социальный превью-текст.
func WithSocietyPreviewText(societyPreviewText string) Option {
	return func(b *Blog) {
		b.societyPreviewText = societyPreviewText
	}
}

// WithAdditionalSEOText устанавливает дополнительный SEO-текст.
func WithAdditionalSEOText(additionalSEOText string) Option {
	return func(b *Blog) {
		b.additionalSEOText = additionalSEOText
	}
}

// WithOrderingNumber устанавливает номер для сортировки.
func WithOrderingNumber(orderingNumber int64) Option {
	return func(b *Blog) {
		b.orderingNumber = orderingNumber
	}
}

// WithCreatedAt устанавливает время создания.
func WithCreatedAt(createdAt time.Time) Option {
	return func(b *Blog) {
		b.createdAt = createdAt
	}
}

// WithIsActive устанавливает статус активности.
func WithIsActive(isActive bool) Option {
	return func(b *Blog) {
		b.isActive = isActive
	}
}
