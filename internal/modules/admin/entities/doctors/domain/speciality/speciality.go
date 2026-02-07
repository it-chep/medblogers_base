package speciality

import "github.com/samber/lo"

// SpecialityID - ID специальности
type SpecialityID int64

// Speciality - справочик специальностей
type Speciality struct {
	id                  SpecialityID
	name                string
	doctorsCount        int64
	isOnlyAdditional    *bool
	primarySpecialities []int64
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

func (c *Speciality) DoctorsCount() int64 {
	return c.doctorsCount
}

// IsAdditional является ли основной специальностью
func (c *Speciality) IsAdditional() bool {
	return len(c.primarySpecialities) > 0 || lo.FromPtr[bool](c.isOnlyAdditional)
}

// PrimaryIDs id основных к этой специальности
func (c *Speciality) PrimaryIDs() []int64 {
	return c.primarySpecialities
}
