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
