package speciality

// SpecialityID - ID специальности
type SpecialityID int64

// Speciality - справочик специальностей
type Speciality struct {
	id   SpecialityID
	name string
}
