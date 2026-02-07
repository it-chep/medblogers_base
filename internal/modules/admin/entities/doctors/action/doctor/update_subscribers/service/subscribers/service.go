package subscribers

import (
	"context"
	"medblogers_base/internal/modules/admin/client/subscribers/indto"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update_subscribers/dto"
)

type Subscribers interface {
	UpdateSubscribers(ctx context.Context, doctorID int64, subscriberItems []indto.UpdateSubscribersItem) error
}

type Service struct {
	subscribers Subscribers
}

func New(subscribers Subscribers) *Service {
	return &Service{
		subscribers: subscribers,
	}
}

func (s *Service) UpdateSubscribers(ctx context.Context, doctorID int64, req dto.UpdateSubscribersRequest) error {
	return s.subscribers.UpdateSubscribers(ctx, doctorID, req.ToInDTO())
}
