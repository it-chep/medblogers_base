package v1

import (
	"medblogers_base/internal/modules/freelancers"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
)

type Implementation struct {
	desc.UnimplementedSeoServer

	freelancers *freelancers.Module
}

// NewFreelancersService return new instance of Implementation.
func NewFreelancersService(freelancers *freelancers.Module) *Implementation {
	return &Implementation{
		freelancers: freelancers,
	}
}
