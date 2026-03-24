package dto

type CreateRequest struct {
	CooperationTypeID    int64
	TopicID              int64
	Title                string
	Description          string
	Price                int64
	ContentFormatID      int64
	BrandID              int64
	PublicationDate      string
	AdMarkingResponsible string
	ResponsesCapacity    int64
	SocialNetworkIDs     []int64
}
