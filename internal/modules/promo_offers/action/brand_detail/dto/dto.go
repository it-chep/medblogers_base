package dto

type Topic struct {
	ID   int64
	Name string
}

type SocialNetwork struct {
	ID   int64
	Name string
	Slug string
	Link string
}

type Brand struct {
	ID             int64
	Title          string
	Slug           string
	Photo          string
	Topic          *Topic
	Website        string
	Description    string
	SocialNetworks []SocialNetwork
}
