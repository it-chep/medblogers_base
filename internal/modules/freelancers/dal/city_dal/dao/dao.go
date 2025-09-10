package dao

import "medblogers_base/internal/modules/freelancers/domain/city"

type CityDAO struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (c CityDAO) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(c.ID),
		city.WithName(c.Name),
	)
}

type CityDAOWithFreelancersCount struct {
	CityDAO
	FreelancersCount int64 `db:"freelancers_count" json:"freelancers_count"`
}

func (c CityDAOWithFreelancersCount) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(c.ID),
		city.WithName(c.Name),
		city.WithFreelancersCount(c.FreelancersCount),
	)
}

type CityDAOWithFreelancerID struct {
	CityDAO
	FreelancerID int64 `db:"freelancer_id" json:"freelancer_id"`
}

func (c CityDAOWithFreelancerID) ToDomain() *city.City {
	return city.BuildCity(
		city.WithID(c.ID),
		city.WithName(c.Name),
		city.WithFreelancersCount(c.FreelancersCount),
	)
}
