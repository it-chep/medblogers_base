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

func (c *City) ID() CityID {
	return c.id
}

func (c *City) Name() string {
	return c.name
}

func (c *City) DoctorsCount() int64 {
	return c.doctorsCount
}
