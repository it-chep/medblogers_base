package entities

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/admin/client"
	blog_action "medblogers_base/internal/modules/admin/entities/blog/action"
	doctor_city_action "medblogers_base/internal/modules/admin/entities/doctors/action/city"
	doctor_action "medblogers_base/internal/modules/admin/entities/doctors/action/doctor"
	doctor_speciality_action "medblogers_base/internal/modules/admin/entities/doctors/action/speciality"

	freelancer_city_action "medblogers_base/internal/modules/admin/entities/freelancers/action/city"
	"medblogers_base/internal/pkg/http"
	"medblogers_base/internal/pkg/postgres"
)

type DoctorModule struct {
	DoctorAgg     *doctor_action.DoctorModuleAggregator
	CityAgg       *doctor_city_action.DoctorCityAggregator
	SpecialityAgg *doctor_speciality_action.DoctorSpecialityAggregator
}

type FreelancerModule struct {
	//DoctorAgg     *freelancer_action.DoctorModuleAggregator
	CityAgg *freelancer_city_action.FreelancerCityAggregator
	//SpecialityAgg *freelancer_speciality_action.DoctorSpecialityAggregator
}

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	BlogModule       *blog_action.BlogModuleAggregator
	DoctorModule     DoctorModule
	FreelancerModule FreelancerModule
}

func NewAggregator(httpConns map[string]http.Executor, config config.AppConfig, pool postgres.PoolWrapper) *Aggregator {

	clients := client.NewAggregator(httpConns, config)

	return &Aggregator{
		BlogModule: blog_action.New(pool, clients),
		DoctorModule: DoctorModule{
			DoctorAgg:     doctor_action.NewDoctorModuleAggregator(clients, pool),
			CityAgg:       doctor_city_action.New(pool),
			SpecialityAgg: doctor_speciality_action.New(pool),
		},
		FreelancerModule: FreelancerModule{
			CityAgg: freelancer_city_action.New(pool),
		},
	}
}
