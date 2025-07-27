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
	//	Если фильтр не подписчиков. Считаем только по городам и специальностям
	if filter.MinSubscribers == 0 && filter.MaxSubscribers == 0 && len(filter.SocialMedia) == 0 {
		doctorsCount, err := a.doctorService.GetDoctorsByFilter(ctx, filter)
		if err != nil {
			logger.Error(ctx, "ошибка при подсчете докторов из базы", err)
			return 0, err
		}
		logger.Message(ctx, fmt.Sprintf("Количество докторов: %d", doctorsCount))
		return doctorsCount, nil
	}

	doctorsIDsFromSubscribers, err := a.subscribersService.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		// если нам приходит 0 или ошибка от сервиса подписчиков, а фильтр был направлен на подписчиков,
		// то мы по честному отдаем что не смогли найти
		logger.Error(ctx, "ошибка при подсчете докторов из подписчиков", err)
		return 0, err
	}

	doctorsCount, err := a.doctorService.GetDoctorsByFilterAndIDs(ctx, filter, doctorsIDsFromSubscribers)
	if err != nil {
		logger.Error(ctx, "ошибка при подсчете докторов из базы", err)
	}

	logger.Message(ctx, fmt.Sprintf("Количество докторов: %d", doctorsCount))
	return doctorsCount, nil
}
