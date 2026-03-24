package dto

type SocialNetworkInput struct {
	SocialNetworkID int64
	Link            string
}

type UpdateRequest struct {
	Photo          string
	Title          string
	Slug           string
	TopicID        int64
	Website        string
	Description    string
	SocialNetworks []SocialNetworkInput
}
