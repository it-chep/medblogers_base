package v1

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) GetPreliminaryFilterCount(ctx context.Context, request *desc.PreliminaryFilterCountRequest) (*desc.PreliminaryFilterCountResponse, error) {
	filter := i.requestToPreliminaryFilterDTO(request)

	freelancersCount, err := i.freelancers.Actions.PreliminaryFilterCount.Do(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &desc.PreliminaryFilterCountResponse{
		FreelancersCount: freelancersCount,
	}, nil
}

func (i *Implementation) requestToPreliminaryFilterDTO(req *desc.PreliminaryFilterCountRequest) freelancer.Filter {
	filter := freelancer.Filter{
		Cities:         req.Cities,
		Specialities:   req.Specialities,
		SocialNetworks: req.Societies,
		PriceCategory:  req.PriceCategory,
	}

	return filter
}
