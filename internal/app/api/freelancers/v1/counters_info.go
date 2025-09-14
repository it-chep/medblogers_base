package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/freelancers/v1"
)

func (i *Implementation) GetCounters(ctx context.Context, _ *desc.GetCountersRequest) (*desc.GetCountersResponse, error) {
	freelancersCount, err := i.freelancers.Actions.GetCounters.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetCountersResponse{
		FreelancersCount: freelancersCount,
	}, nil
}
