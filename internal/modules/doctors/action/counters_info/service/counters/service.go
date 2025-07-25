package counters

import (
	"context"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
)

// DoctorsStorage .
type DoctorsStorage interface {
	GetDoctorsCount(ctx context.Context) (int64, error)
}

type SubscribersGetter interface {
	GetAllSubscribersInfo(ctx context.Context) (indto.GetAllSubscribersInfoResponse, error)
}

type Service struct {
	store             DoctorsStorage
	subscribersGetter SubscribersGetter
}

func NewService(doctorsStorage DoctorsStorage, subscribersGetter SubscribersGetter) *Service {
	return &Service{
		store:             doctorsStorage,
		subscribersGetter: subscribersGetter,
	}
}

// GetDoctorsCount получение общего количества врачей
func (s *Service) GetDoctorsCount(ctx context.Context) (int64, error) {
	return s.store.GetDoctorsCount(ctx)
}

// GetSubscribersCount получение общего количества подписчиков всех врачей
func (s *Service) GetSubscribersCount(ctx context.Context) (string, error) {
	subs, err := s.subscribersGetter.GetAllSubscribersInfo(ctx)
	if err != nil {
		return "", err
	}

	return subs.SubscribersCount, nil
}
