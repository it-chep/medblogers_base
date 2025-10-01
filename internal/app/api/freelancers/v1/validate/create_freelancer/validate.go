package create_freelancer

type CreateFreelancerRequest struct {
	Email            string `json:"email" validate:"required,email"`
	LastName         string `json:"lastName" validate:"required,max=100"`
	FirstName        string `json:"firstName" validate:"required,max=100"`
	MiddleName       string `json:"middleName" validate:"required,max=100"`
	TelegramUsername string `json:"telegramUsername" validate:"required,max=100"`

	WorkingExperience int64 `json:"workingExperience" validate:"required,max=100"`
	AgreePolicy       bool  `json:"agreePolicy" validate:"required"`

	CityID       int64 `json:"cityID" validate:"required"`
	SpecialityID int64 `json:"specialityID" validate:"required"`
}

type ValidationError struct {
	Code  int    `json:"code"`
	Text  string `json:"text"`
	Field string `json:"field"`
}
