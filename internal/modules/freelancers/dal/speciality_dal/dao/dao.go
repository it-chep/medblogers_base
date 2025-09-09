package dao

import "medblogers_base/internal/modules/freelancers/domain/speciality"

type SpecialityDAO struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s SpecialityDAO) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(s.ID),
		speciality.WithName(s.Name),
	)
}

type SpecialityDAOWithFreelancersCount struct {
	SpecialityDAO
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}

func (s SpecialityDAOWithFreelancersCount) ToDomain() *speciality.Speciality {
	return speciality.BuildSpeciality(
		speciality.WithID(s.ID),
		speciality.WithName(s.Name),
		speciality.WithFreelancersCount(s.FreelancersCount),
	)
}
