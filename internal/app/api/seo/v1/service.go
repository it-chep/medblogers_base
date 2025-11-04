package v1

import (
	"medblogers_base/internal/modules/doctors"
	"medblogers_base/internal/modules/freelancers"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
)

type Implementation struct {
	desc.UnimplementedSeoServer

	doctors    *doctors.Module
	freelancer *freelancers.Module
}

// NewSeoService return new instance of Implementation.
func NewSeoService(doctors *doctors.Module, freelancer *freelancers.Module) *Implementation {
	return &Implementation{
		doctors:    doctors,
		freelancer: freelancer,
	}
}
