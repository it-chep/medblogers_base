package v1

import (
	"medblogers_base/internal/modules/doctors"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
	"medblogers_base/internal/pkg/config"
)

type Implementation struct {
	desc.UnimplementedDoctorServiceServer

	mutableConfig config.Config
	doctors       *doctors.Module
}

// NewDoctorsService return new instance of Implementation.
func NewDoctorsService(doctors *doctors.Module, mutableConfig config.Config) *Implementation {
	return &Implementation{
		doctors:       doctors,
		mutableConfig: mutableConfig,
	}
}
