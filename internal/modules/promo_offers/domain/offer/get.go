package offer

import (
	"time"

	"github.com/google/uuid"
)

func (o *Offer) GetID() uuid.UUID {
	return o.id
}

func (o *Offer) GetCooperationTypeID() int64 {
	return o.cooperationTypeID
}

func (o *Offer) GetBusinessCategoryID() int64 {
	return o.businessCategoryID
}

func (o *Offer) GetTitle() string {
	return o.title
}

func (o *Offer) GetDescription() string {
	return o.description
}

func (o *Offer) GetPrice() int64 {
	return o.price
}

func (o *Offer) GetContentFormatID() int64 {
	return o.contentFormatID
}

func (o *Offer) GetBrandID() int64 {
	return o.brandID
}

func (o *Offer) GetPublicationDate() *time.Time {
	return o.publicationDate
}

func (o *Offer) GetAdMarkingResponsible() string {
	return o.adMarkingResponsible
}

func (o *Offer) GetResponsesCapacity() int64 {
	return o.responsesCapacity
}

func (o *Offer) GetIsActive() bool {
	return o.isActive
}

func (o *Offer) GetCreatedAt() time.Time {
	return o.createdAt
}
