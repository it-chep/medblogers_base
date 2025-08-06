package dto

// Абстракция над HTTP

type CreateDoctorRequest struct {
	// username telegram
	Telegram string `json:"telegram"`
	// username telegram
	Instagram string `json:"instagram"`
}

type UpdateDoctorRequest struct {
	// username telegram
	Telegram string `json:"telegram"`
	// username telegram
	Instagram string `json:"instagram"`
}

type DoctorsFilterWithIDsRequest struct {
	SocialMedia []string `json:"social_media"`

	MaxSubscribers int64 `json:"max_subscribers"`
	MinSubscribers int64 `json:"min_subscribers"`

	Limit       int64  `json:"limit"`
	CurrentPage int64  `json:"current_page"`
	Sort        string `json:"sort"`

	DoctorIDs []int64 `json:"doctor_ids"`
}
