package dto

// Абстракция над HTTP

type CreateDoctorRequest struct {
	DoctorID int64 `json:"doctor_id"`
	// username
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Youtube   string `json:"youtube"`
}

type UpdateDoctorRequest struct {
	// username telegram
	Telegram string `json:"telegram"`
	// username telegram
	Instagram string `json:"instagram"`

	IsActive bool `json:"is_active"`
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

type SubscribersItem struct {
	Key       string `json:"key"`
	SubsCount int64  `json:"subs_count"`
}

type UpdateSubscribersRequest struct {
	Items []SubscribersItem `json:"items"`
}

type ChangeVipActivityRequest struct {
	Activity bool `json:"vip_activity"`
}
