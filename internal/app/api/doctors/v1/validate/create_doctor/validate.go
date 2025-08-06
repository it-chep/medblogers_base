package create_doctor

type CreateDoctorRequest struct {
	Email            string `json:"email" validate:"required,email"`
	LastName         string `json:"last_name" validate:"required,max=100"`
	FirstName        string `json:"first_name" validate:"required,max=100"`
	MiddleName       string `json:"middle_name" validate:"required,max=100"`
	BirthDate        string `json:"birth_date" validate:"required"`
	TelegramUsername string `json:"telegram_username" validate:"required,max=100"`

	AgreePolicy bool `json:"agree_policy" validate:"required"`

	CityID       int64 `json:"city_id" validate:"required"`
	SpecialityID int64 `json:"speciality_id" validate:"required"`
}

type ValidationError struct {
	Code  int    `json:"code"`
	Text  string `json:"text"`
	Field string `json:"field"`
}
