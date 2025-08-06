package dto

type DoctorData struct {
	Name                 string `json:"name"`
	Age                  int    `json:"age"`
	Birthday             string `json:"birthday"`
	InstagramURL         string `json:"inst_url"`
	VKURL                string `json:"vk_url"`
	DzenURL              string `json:"dzen_url"`
	TelegramURL          string `json:"tg_url"`
	Subscribers          int    `json:"subscribers"`
	City                 string `json:"city"`
	MedicalDirections    string `json:"medical_directions"`
	Speciality           string `json:"speciallity"`
	AdditionalSpeciality string `json:"additional_speciallity"`
	MainTheme            string `json:"main_theme"`
}
