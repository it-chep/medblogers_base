package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
)

// GetPagesCount - /api/v1/pages_count [GET]
func (i *Implementation) GetPagesCount(ctx context.Context, req *desc.PagesCountRequest) (*desc.PagesCountResponse, error) {
	filter := i.requestToPagesCountDTO(req)

	pagesCount, err := i.doctors.Actions.GetPagesCount.Do(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &desc.PagesCountResponse{
		PagesCount: pagesCount,
	}, nil
}

func (i *Implementation) requestToPagesCountDTO(req *desc.PagesCountRequest) dto.Filter {
	maxSubscribers := req.MaxSubscribers
	if maxSubscribers <= 0 {
		maxSubscribers = 5_000_000
	}

	minSubscribers := req.MinSubscribers
	if minSubscribers <= 0 {
		minSubscribers = 100
	}

	return dto.Filter{
		MaxSubscribers: maxSubscribers,
		MinSubscribers: minSubscribers,
		Cities:         req.Cities,
		Specialities:   req.Specialities,
		SocialMedia:    req.SocialMedia,
	}
}
