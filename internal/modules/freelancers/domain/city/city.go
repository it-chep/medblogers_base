package city

// City - справочик городов
type City struct {
	id               int64
	name             string
	freelancersCount int64
}

// BuildCity создать город
func BuildCity(options ...Option) *City {
	e := &City{}
	for _, option := range options {
		option(e)
	}
	return e
}

func (c *City) ID() int64 {
	return c.id
}

func (c *City) Name() string {
	return c.name
}

func (c *City) FreelancersCount() int64 {
	return c.freelancersCount
}
