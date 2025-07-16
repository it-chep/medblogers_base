package doctor_detail

import (
	"context"
	"fmt"
	"medblogers_base/internal/pkg/logger"

	"medblogers_base/internal/modules/doctors/action/doctor_detail/dal"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/dto"
	"medblogers_base/internal/modules/doctors/action/doctor_detail/service/enricher"
	"medblogers_base/internal/modules/doctors/client"
)

// Action получение детальной информации о докторе
type Action struct {
	doctorStorage       *dal.Repository
	subscribersEnricher *enricher.SubscribersEnricher
}

// New .
func New(clients *client.Aggregator) *Action {
	return &Action{
		subscribersEnricher: enricher.NewSubscribersEnricher(clients.Subscribers),
	}
}

func (a Action) Do(ctx context.Context, doctorID int64) (dto.DoctorDTO, error) {
	logger.Message(ctx, fmt.Sprintf("[DoctorDetail] Получение данных о докторе %d", doctorID))
	doc, err := a.doctorStorage.GetDoctorInfo(ctx, doctorID)
	if err != nil {
		return dto.DoctorDTO{}, err // 404 not found
	}

	docDTO := dto.New(doc)
	docDTO, err = a.subscribersEnricher.Enrich(ctx, doc.GetID(), docDTO)
	if err != nil {
		// логируем ошибку от подписчиков
		return docDTO, nil
	}

	// todo мб параллель
	//	обогатить доп специальностями
	//	обогатить доп городами
	//	обогатить подписчиками

	return docDTO, nil
}
