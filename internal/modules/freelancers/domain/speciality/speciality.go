package speciality

// Speciality - справочик специальностей
type Speciality struct {
	id               int64
	name             string
	freelancersCount int64
}

type Specialities []*Speciality

// BuildSpeciality создать специальность
func BuildSpeciality(options ...Option) *Speciality {
	e := &Speciality{}
	for _, option := range options {
		option(e)
	}
	return e
}

func (c *Speciality) ID() int64 {
	return c.id
}

func (c *Speciality) Name() string {
	return c.name
}

func (c *Speciality) FreelancersCount() int64 {
	return c.freelancersCount
}
