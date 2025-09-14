package v1

import (
	"context"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) GetPagesCount(ctx context.Context, request *desc.PagesCountRequest) (*desc.PagesCountResponse, error) {
	filter := i.requestToPagesCountDTO(request)

	pagesCount, err := i.freelancers.Actions.GetPagesCount.Do(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &desc.PagesCountResponse{
		PagesCount: pagesCount,
	}, nil
}

func (i *Implementation) requestToPagesCountDTO(req *desc.PagesCountRequest) freelancer.Filter {
	filter := freelancer.Filter{
		Cities:         req.Cities,
		Specialities:   req.Specialities,
		SocialNetworks: req.Societies,
		PriceCategory:  req.PriceCategory,
	}

	if req.ExperienceWithDoctors != nil {
		filter.ExperienceWithDoctors = req.ExperienceWithDoctors
	}

	return filter
}
