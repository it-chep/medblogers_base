package dto

const (
	CreateDoctorEvent = "doctor_on_site_notification"
)

type CreateDoctorNotification struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	InstURL  string `json:"inst_url"`
	TgURL    string `json:"tg_url"`
	Message  string `json:"message"`
	ClientID int64  `json:"client_id"`
}
