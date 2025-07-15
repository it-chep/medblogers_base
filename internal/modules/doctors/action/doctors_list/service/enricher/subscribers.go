package enricher

import (
	"context"

	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SubscribersGetter

// SubscribersGetter получение информации о подписчиках
type SubscribersGetter interface {
	GetSubscribersByDoctorIDs(ctx context.Context, medblogersIDs doctor.MedblogersIDs) (map[doctor.MedblogersID]indto.GetSubscribersByDoctorIDsResponse, error)
}

type Service struct {
	getter SubscribersGetter
}

func New(getter SubscribersGetter) *Service {
	return &Service{
		getter: getter,
	}
}

// EnrichSubscribers - обогащение подписчиками в миниатюры докторов
func (s *Service) EnrichSubscribers(ctx context.Context, doctorsIDs doctor.MedblogersIDs, docDTO dto.DoctorDTO) error {
	_, err := s.getter.GetSubscribersByDoctorIDs(ctx, doctorsIDs)
	if err != nil {
		return err
	}

	return nil
}
