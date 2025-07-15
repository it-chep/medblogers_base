package dao

import "medblogers_base/internal/modules/doctors/domain/doctor"

// DoctorDAO .
type DoctorDAO struct {
}

// ToDomain конвертирует объект доступа к данным в доменное представление
func (d DoctorDAO) ToDomain() *doctor.Doctor {
	return &doctor.Doctor{}
}
