package v1

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
)

// GetPreliminaryFilterCount - /api/v1/preliminary_filter_count [GET]
func (i *Implementation) GetPreliminaryFilterCount(ctx context.Context, req *desc.PreliminaryFilterCountRequest) (*desc.PreliminaryFilterCountResponse, error) {
	filter := i.requestToPreliminaryFilterDTO(req)

	doctorsCount, err := i.doctors.Actions.PreliminaryFilterCount.Do(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &desc.PreliminaryFilterCountResponse{
		DoctorsCount: doctorsCount,
	}, nil
}

func (i *Implementation) requestToPreliminaryFilterDTO(req *desc.PreliminaryFilterCountRequest) dto.Filter {
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
