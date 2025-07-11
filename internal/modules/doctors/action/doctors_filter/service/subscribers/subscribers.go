package subscribers

import (
	"context"

	"github.com/it-chep/medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/doctor"
)

// SubscribersGetter получение информации о подписчиках
type SubscribersGetter interface {
	GetDoctorsByFilter(ctx context.Context, request indto.GetDoctorsByFilterRequest) ([]indto.GetDoctorsByFilterResponse, error)
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
	ds, err := s.getter.GetSubscribersByDoctorIDs(ctx, doctorsIDs)
	if err != nil {
		return err
	}
}

// FilterDoctorsBySubscribers - фильтрация по количеству подписчиков
func (s *Service) FilterDoctorsBySubscribers(ctx context.Context) {
	filter, err := s.getter.GetDoctorsByFilter(ctx, fil)
	if err != nil {
		return
	}
}
