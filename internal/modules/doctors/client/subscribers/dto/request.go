package dto

// Абстракция над HTTP

type CreateDoctorRequest struct {
	DoctorID int64 `json:"doctor_id"`
	// username
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Youtube   string `json:"youtube"`
	Vk        string `json:"vk"`
}

type UpdateDoctorRequest struct {
	// username telegram
	Telegram string `json:"telegram"`
	// username telegram
	Instagram string `json:"instagram"`
}

type DoctorsFilterWithIDsRequest struct {
	SocialMedia []string `json:"social_media,omitempty"`

	MaxSubscribers int64 `json:"max_subscribers,omitempty"`
	MinSubscribers int64 `json:"min_subscribers,omitempty"`

	Limit       int64  `json:"limit,omitempty"`
	CurrentPage int64  `json:"current_page,omitempty"`
	Sort        string `json:"sort,omitempty"`

	DoctorIDs []int64 `json:"doctor_ids,omitempty"`
}

type CheckTelegramInBlackListRequest struct {
	Telegram string `json:"telegram"`
}
