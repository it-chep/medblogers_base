package create_doctor

type CreateDoctorRequest struct {
	Email                 string  `json:"email" validate:"required,email"`
	LastName              string  `json:"last_name" validate:"required,max=100"`
	FirstName             string  `json:"first_name" validate:"required,max=100"`
	MiddleName            string  `json:"middle_name" validate:"required,max=100"`
	BirthDate             string  `json:"birth_date" validate:"required"`
	AdditionalCities      []int64 `json:"additional_cities"`
	AdditionalSpecialties []int   `json:"additional_specialties"`
	InstagramUsername     string  `json:"instagram_username"`
	VKUsername            string  `json:"vk_username"`
	TelegramUsername      string  `json:"telegram_username" validate:"required"`
	DzenUsername          string  `json:"dzen_username"`
	YoutubeUsername       string  `json:"youtube_username"`
	TelegramChannel       string  `json:"telegram_channel"`
	CityID                int64   `json:"city_id" validate:"required"`
	SpecialityID          int64   `json:"speciality_id" validate:"required"`
	MainBlogTheme         string  `json:"main_blog_theme"`
	SiteLink              string  `json:"site_link"`
	AgreePolicy           bool    `json:"agree_policy" validate:"required"`
}

type ValidationErrors struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Response struct {
	Errors []ValidationErrors `json:"errors"`
}
