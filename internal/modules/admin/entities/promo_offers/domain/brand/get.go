package brand

import "time"

func (b *Brand) GetID() int64 {
	return b.id
}

func (b *Brand) GetPhoto() string {
	return b.photo
}

func (b *Brand) GetTitle() string {
	return b.title
}

func (b *Brand) GetSlug() string {
	return b.slug
}

func (b *Brand) GetBusinessCategoryID() int64 {
	return b.businessCategoryID
}

func (b *Brand) GetWebsite() string {
	return b.website
}

func (b *Brand) GetDescription() string {
	return b.description
}

func (b *Brand) GetAbout() string {
	return b.about
}

func (b *Brand) GetIsActive() bool {
	return b.isActive
}

func (b *Brand) GetCreatedAt() *time.Time {
	return b.createdAt
}
