package dto

const (
	CreateFreelancerEvent = "freelancer_on_site_notification"
)

type CreateFreelancerNotification struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	TgURL    string `json:"tg_url"`
	Message  string `json:"message"`
	ClientID int64  `json:"client_id"`
}
