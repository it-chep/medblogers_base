package subscribers

import "context"

type Subscribers interface {
	ChangeVipActivity(ctx context.Context, doctorID int64, activity bool) error
}

type Service struct {
	subscribers Subscribers
}

func New(subscribers Subscribers) *Service {
	return &Service{
		subscribers: subscribers,
	}
}

func (s *Service) ChangeVipActivity(ctx context.Context, doctorID int64, activity bool) error {
	return s.subscribers.ChangeVipActivity(ctx, doctorID, activity)
}
