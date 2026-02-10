package dto

type UpdateRequest struct {
	Name          string
	Slug          string
	PortfolioLink string
	TgURL         string

	MainCityID       int64
	MainSpecialityID int64

	AgencyRepresentative bool

	DateStarted       string
	CooperationTypeID int64
	PriceCategory     int64
}
