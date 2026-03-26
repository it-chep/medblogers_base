package brand

import "time"

type Option func(*Brand)

func WithID(id int64) Option {
	return func(b *Brand) {
		b.id = id
	}
}

func WithPhoto(photo string) Option {
	return func(b *Brand) {
		b.photo = photo
	}
}

func WithTitle(title string) Option {
	return func(b *Brand) {
		b.title = title
	}
}

func WithSlug(slug string) Option {
	return func(b *Brand) {
		b.slug = slug
	}
}

func WithBusinessCategoryID(businessCategoryID int64) Option {
	return func(b *Brand) {
		b.businessCategoryID = businessCategoryID
	}
}

func WithWebsite(website string) Option {
	return func(b *Brand) {
		b.website = website
	}
}

func WithDescription(description string) Option {
	return func(b *Brand) {
		b.description = description
	}
}

func WithIsActive(isActive bool) Option {
	return func(b *Brand) {
		b.isActive = isActive
	}
}

func WithCreatedAt(createdAt *time.Time) Option {
	return func(b *Brand) {
		b.createdAt = createdAt
	}
}
