package doctor

import (
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/activate"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/add_additional_city"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/add_additional_speciality"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/deactivate"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/delete_additional_city"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/delete_additional_speciality"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_by_id"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_cooperation_types"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_doctor_additional_cities"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/get_doctor_additional_specialities"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/save_doctor_photo"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update_subscribers"
	"medblogers_base/internal/pkg/postgres"
)

type DoctorModuleAggregator struct {
	GetDoctors *get.Action

	GetDoctorByID *get_by_id.Action

	ActivateDoctor   *activate.Action
	DeactivateDoctor *deactivate.Action

	AddAdditionalCity    *add_additional_city.Action
	DeleteAdditionalCity *delete_additional_city.Action

	AddAdditionalSpeciality    *add_additional_speciality.Action
	DeleteAdditionalSpeciality *delete_additional_speciality.Action

	SaveDoctorPhoto   *save_doctor_photo.Action
	UpdateSubscribers *update_subscribers.Action
	UpdateDoctor      *update.Action

	GetCooperationTypes             *get_cooperation_types.Action
	GetDoctorAdditionalCities       *get_doctor_additional_cities.Action
	GetDoctorAdditionalSpecialities *get_doctor_additional_specialities.Action
}

func NewDoctorModuleAggregator(clients *client.Aggregator, pool postgres.PoolWrapper) *DoctorModuleAggregator {
	return &DoctorModuleAggregator{
		GetDoctors: get.New(pool),

		GetDoctorByID: get_by_id.New(clients, pool),

		ActivateDoctor:   activate.New(clients, pool),
		DeactivateDoctor: deactivate.New(clients, pool),

		AddAdditionalCity:    add_additional_city.New(pool),
		DeleteAdditionalCity: delete_additional_city.New(pool),

		AddAdditionalSpeciality:    add_additional_speciality.New(pool),
		DeleteAdditionalSpeciality: delete_additional_speciality.New(pool),

		SaveDoctorPhoto:   save_doctor_photo.New(clients, pool),
		UpdateSubscribers: update_subscribers.New(clients),
		UpdateDoctor:      update.New(pool),

		GetCooperationTypes:             get_cooperation_types.New(pool),
		GetDoctorAdditionalCities:       get_doctor_additional_cities.New(pool),
		GetDoctorAdditionalSpecialities: get_doctor_additional_specialities.New(pool),
	}
}
