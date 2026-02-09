package dto

type UpdateRequest struct {
	Name  string
	Slug  string
	Email string

	InstURL      string
	VkURL        string
	DzenURL      string
	TgURL        string
	TgChannelURL string
	YouTubeURL   string
	TikTokURL    string
	SiteLink     string

	MainCityID       int64
	MainSpecialityID int64

	MainBlogTheme string
	IsKFDoctor    bool

	BirthDate            string
	CooperationTypeID    int64
	MedicalDirections    string
	MarketingPreferences string
}
