package offer

import (
	"time"

	"github.com/google/uuid"
)

type Option func(*Offer)

func WithID(id uuid.UUID) Option {
	return func(o *Offer) {
		o.id = id
	}
}

func WithCooperationTypeID(cooperationTypeID int64) Option {
	return func(o *Offer) {
		o.cooperationTypeID = cooperationTypeID
	}
}

func WithTopicID(topicID int64) Option {
	return func(o *Offer) {
		o.topicID = topicID
	}
}

func WithTitle(title string) Option {
	return func(o *Offer) {
		o.title = title
	}
}

func WithDescription(description string) Option {
	return func(o *Offer) {
		o.description = description
	}
}

func WithPrice(price int64) Option {
	return func(o *Offer) {
		o.price = price
	}
}

func WithContentFormatID(contentFormatID int64) Option {
	return func(o *Offer) {
		o.contentFormatID = contentFormatID
	}
}

func WithBrandID(brandID int64) Option {
	return func(o *Offer) {
		o.brandID = brandID
	}
}

func WithPublicationDate(publicationDate *time.Time) Option {
	return func(o *Offer) {
		o.publicationDate = publicationDate
	}
}

func WithAdMarkingResponsible(adMarkingResponsible string) Option {
	return func(o *Offer) {
		o.adMarkingResponsible = adMarkingResponsible
	}
}

func WithResponsesCapacity(responsesCapacity int64) Option {
	return func(o *Offer) {
		o.responsesCapacity = responsesCapacity
	}
}

func WithIsActive(isActive bool) Option {
	return func(o *Offer) {
		o.isActive = isActive
	}
}

func WithCreatedAt(createdAt *time.Time) Option {
	return func(o *Offer) {
		o.createdAt = createdAt
	}
}
