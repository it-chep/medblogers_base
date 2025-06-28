package speciality

// SpecialityID - ID специальности
type SpecialityID int64

// Speciality - справочик специальностей
type Speciality struct {
	id           SpecialityID
	name         string
	doctorsCount int64
}

// BuildSpeciality создать специальность
func BuildSpeciality(options ...Option) *Speciality {
	e := &Speciality{}
	for _, option := range options {
		option(e)
	}
	return e
}
