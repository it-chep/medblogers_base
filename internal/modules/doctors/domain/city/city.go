package city

// CityID - ID города
type CityID int64

// City - справочик городов
type City struct {
	id           CityID
	name         string
	doctorsCount int64
}

// BuildCity создать город
func BuildCity(options ...Option) *City {
	e := &City{}
	for _, option := range options {
		option(e)
	}
	return e
}
