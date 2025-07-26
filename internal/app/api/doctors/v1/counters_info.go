package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/doctors/v1"
)

// GetCounters - /api/v1/counters_info [GET]
func (i *Implementation) GetCounters(ctx context.Context, _ *desc.GetCountersRequest) (*desc.GetCountersResponse, error) {
	countersDomain, err := i.doctors.Actions.CounterInfo.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetCountersResponse{
		DoctorsCount:     countersDomain.DoctorsCount,
		SubscribersCount: countersDomain.SubscribersCount,
	}, nil
}
