package city

// CityID - ID города
type CityID int64

// City - справочик городов
type City struct {
	id   CityID
	name string
}
