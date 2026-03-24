package offer

import (
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Offer struct {
	id                   uuid.UUID
	cooperationTypeID    int64
	topicID              int64
	title                string
	description          string
	price                int64
	contentFormatID      int64
	brandID              int64
	publicationDate      *time.Time
	adMarkingResponsible string
	responsesCapacity    int64
	isActive             bool
	createdAt            *time.Time
}

func New(options ...Option) *Offer {
	item := &Offer{}
	for _, option := range options {
		option(item)
	}

	return item
}

type Offers []*Offer

func (o Offers) IDs() []uuid.UUID {
	return lo.Map(o, func(item *Offer, _ int) uuid.UUID {
		return item.GetID()
	})
}

func (o Offers) BrandIDs() []int64 {
	return lo.FilterMap(o, func(item *Offer, _ int) (int64, bool) {
		if item.GetBrandID() <= 0 {
			return 0, false
		}

		return item.GetBrandID(), true
	})
}

func (o Offers) CooperationTypeIDs() []int64 {
	return lo.FilterMap(o, func(item *Offer, _ int) (int64, bool) {
		if item.GetCooperationTypeID() <= 0 {
			return 0, false
		}

		return item.GetCooperationTypeID(), true
	})
}

func (o Offers) TopicIDs() []int64 {
	return lo.FilterMap(o, func(item *Offer, _ int) (int64, bool) {
		if item.GetTopicID() <= 0 {
			return 0, false
		}

		return item.GetTopicID(), true
	})
}

func (o Offers) ContentFormatIDs() []int64 {
	return lo.FilterMap(o, func(item *Offer, _ int) (int64, bool) {
		if item.GetContentFormatID() <= 0 {
			return 0, false
		}

		return item.GetContentFormatID(), true
	})
}
