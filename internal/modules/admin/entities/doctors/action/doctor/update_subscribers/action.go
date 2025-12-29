package update_subscribers

import (
	"context"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update_subscribers/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update_subscribers/service/subscribers"
)

type Action struct {
	subscribers *subscribers.Service
}

func New(clients *client.Aggregator) *Action {
	return &Action{
		subscribers: subscribers.New(clients.Subscribers),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64, req dto.UpdateSubscribersRequest) error {
	return a.subscribers.UpdateSubscribers(ctx, doctorID, req)
}
