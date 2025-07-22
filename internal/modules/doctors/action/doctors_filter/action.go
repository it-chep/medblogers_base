package doctors_filter

import (
	"context"
	"fmt"
	consts "medblogers_base/internal/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dal"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/doctor"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/subscribers"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/samber/lo"
)

const firstPage = 1

// Action фильтрация докторов
type Action struct {
	subscribersFilter *subscribers.Service
	doctorsFilter     *doctor.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		subscribersFilter: subscribers.New(clients.Subscribers),
		doctorsFilter: doctor.New(
			dal.NewRepository(pool),
			dal.NewRepository(pool),
			clients.S3,
			clients.Subscribers,
		),
	}
}

func (a Action) Do(ctx context.Context, filter *dto.Filter) (dto.Response, error) {
	logger.Message(ctx, fmt.Sprintf("[Filter] Фильтрация докторов: %v", filter))
	// Дефолтная загрузка без фильтров
	if filter == nil {
		doctorsMap, err := a.doctorsFilter.GetDoctors(ctx, firstPage)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при получении докторов", err)
			return dto.Response{}, err
		}

		return dto.Response{
			Doctors:     lo.Values(doctorsMap),
			CurrentPage: firstPage,
			Pages:       max(int64(len(doctorsMap)), consts.LimitDoctorsOnPage),
		}, nil
	}

	//	Если фильтр не подписчиков. Берем докторов, обогажаем миниатюрами
	if filter.MinSubscribers == 0 || filter.MaxSubscribers == 0 || len(filter.SocialMedia) == 0 {
		doctorsMap, err := a.doctorsFilter.GetDoctorsByFilter(ctx, filter)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при получении докторов", err)
			return dto.Response{}, err
		}

		return dto.Response{
			Doctors:     lo.Values(doctorsMap),
			CurrentPage: filter.Page,
			Pages:       max(int64(len(doctorsMap)), consts.LimitDoctorsOnPage),
		}, nil
	}

	logger.Message(ctx, "[Filter] Фильтруем по подписчикам")
	subsResponse, err := a.subscribersFilter.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[Filter] Ошибка при фильтрации докторов по подписчиками, делаем дефолт", err)
	}
	doctorsMap, err := a.doctorsFilter.GetDoctorsByIDs(ctx, filter.Page, lo.Keys(subsResponse))
	if err != nil {
		logger.Error(ctx, "[Filter] Ошибка при обогащении фотографией", err)
	}

	// метчим 2 мапы с разными данными
	a.subscribersFilter.EnrichSubscribers(doctorsMap, subsResponse)

	return dto.Response{
		Doctors:     lo.Values(doctorsMap),
		CurrentPage: filter.Page,
		Pages:       max(int64(len(doctorsMap)), consts.LimitDoctorsOnPage),
	}, nil
}
