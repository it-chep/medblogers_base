package enricher

import (
	"context"

	"github.com/it-chep/medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/domain/doctor"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . SubscribersGetter

type SubscribersGetter interface {
	GetDoctorSubscribers(ctx context.Context, medblogersID doctor.MedblogersID) (indto.GetDoctorSubscribersResponse, error)
}

type SubscribersEnricher struct {
	getter SubscribersGetter
}

func NewSubscribersEnricher(getter SubscribersGetter) *SubscribersEnricher {
	return &SubscribersEnricher{
		getter: getter,
	}
}

// Enrich - обогащение подписчиками
func (e *SubscribersEnricher) Enrich(ctx context.Context, doctorID doctor.MedblogersID, docDTO dto.DoctorDTO) (dto.DoctorDTO, error) {

	subscribersInfo, err := e.getter.GetDoctorSubscribers(ctx, doctorID)
	if err != nil {
		return docDTO, err
	}

	docDTO.TgSubsCount = subscribersInfo.TgSubsCount
	docDTO.TgSubsCountText = subscribersInfo.TgSubsCountText
	docDTO.TgLastUpdatedDate = subscribersInfo.TgLastUpdatedDate
	docDTO.InstSubsCount = subscribersInfo.InstSubsCount
	docDTO.InstSubsCountText = subscribersInfo.InstSubsCountText
	docDTO.InstLastUpdatedDate = subscribersInfo.InstLastUpdatedDate

	return docDTO, nil
}
