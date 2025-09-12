package dto

type PriceListItem struct {
	Name  string
	Price int64
}

type PriceList []PriceListItem

type CreateRequest struct {
	ID int64

	FirstName  string
	LastName   string
	MiddleName string

	Email                    string
	Slug                     string
	Name                     string
	HasExperienceWithDoctors bool
	TgUsername               string
	PortfolioLink            string
	MainSpecialityID         int64
	MainCityID               int64
	PriceCategory            int64

	AdditionalCities      []int64
	AdditionalSpecialties []int64

	PriceList PriceList
}

type ValidationError struct {
	Code  int
	Text  string
	Field string
}

func (e ValidationError) Error() string {
	return e.Text
}
