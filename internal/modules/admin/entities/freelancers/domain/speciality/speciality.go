package speciality

// SpecialityID - ID специальности
type SpecialityID int64

// Speciality - справочик специальностей
type Speciality struct {
	id   SpecialityID
	name string
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

func (c *Speciality) ID() SpecialityID {
	return c.id
}

func (c *Speciality) Name() string {
	return c.name
}
