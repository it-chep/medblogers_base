package preliminary_filter_count

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dal"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/service/doctor"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/service/subscribers"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	doctorService      *doctor.Service
	subscribersService *subscribers.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		doctorService:      doctor.New(dal.NewRepository(pool)),
		subscribersService: subscribers.New(clients.Subscribers),
	}
}

func (a *Action) Do(ctx context.Context, filter dto.Filter) (int64, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[PreliminaryFilterCount] Предфильтрация докторов: MaxSubs: %d, MinSubs: %d, SocialMedia: %v, Cities: %v, Specialities: %v, Page: %d",
		filter.MinSubscribers, filter.MinSubscribers, filter.SocialMedia, filter.Cities, filter.Specialities, filter.Page,
	))

	// Получение по фильтрам из базы если есть
	if len(filter.Cities) != 0 || len(filter.Specialities) != 0 {
		// Получаем всех докторов без лимита, так как у нас стоят индексы и идем сокращать выборку в подписчиков
		doctorIDs, err := a.doctorService.GetDoctorsByFilter(ctx, filter)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при получении докторов", err)
			return 0, err
		}

		// Фильтруем и сортируем в подписчиках
		doctorsCount, err := a.subscribersService.FilterDoctorsBySubscribersWithDoctorsIDs(ctx, filter, doctorIDs)
		if err != nil {
			return 0, err
		}

		logger.Message(ctx, fmt.Sprintf("Количество докторов: %d", doctorsCount))

		return doctorsCount, nil
	}

	// Фильтр только по подписчикам
	logger.Message(ctx, "[Filter] Фильтруем по подписчикам")
	doctorsCount, err := a.subscribersService.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[Filter] Ошибка при фильтрации докторов по подписчиками, делаем дефолт", err)
	}

	logger.Message(ctx, fmt.Sprintf("Количество докторов: %d", doctorsCount))

	return doctorsCount, nil
}
