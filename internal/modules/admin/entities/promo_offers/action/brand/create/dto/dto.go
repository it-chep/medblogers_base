package dto

type SocialNetworkInput struct {
	SocialNetworkID int64
	Link            string
}

type CreateRequest struct {
	Photo          string
	Title          string
	Slug           string
	TopicID        int64
	Website        string
	Description    string
	SocialNetworks []SocialNetworkInput
}
