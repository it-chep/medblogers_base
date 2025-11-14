package preliminary_filter_count

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dal"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/dto"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/service/doctor"
	"medblogers_base/internal/modules/doctors/action/preliminary_filter_count/service/subscribers"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	doctorService      *doctor.Service
	subscribersService *subscribers.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		doctorService:      doctor.New(dal.NewRepository(pool), doctor_dal.NewRepository(pool)),
		subscribersService: subscribers.New(clients.Subscribers),
	}
}

func (a *Action) Do(ctx context.Context, filter dto.Filter) (int64, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[PreliminaryFilterCount] Предфильтрация докторов: MaxSubs: %d, MinSubs: %d, SocialMedia: %v, Cities: %v, Specialities: %v, Page: %d",
		filter.MinSubscribers, filter.MinSubscribers, filter.SocialMedia, filter.Cities, filter.Specialities, filter.Page,
	))

	if a.defaultFilter(filter) {
		return a.fallbackDoctorsOnlySubsFilter(ctx)
	}

	// Получение по фильтрам из базы если есть
	if len(filter.Cities) != 0 || len(filter.Specialities) != 0 {
		return a.getDoctorsByCitiesAndSpecialitiesFilter(ctx, filter)
	}

	// Фильтр только по подписчикам
	logger.Message(ctx, "[Filter] Фильтруем по подписчикам")
	doctorsCount, err := a.subscribersService.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[Error] Ошибка при фильтрации докторов по подписчиками, делаем дефолт", err)
		return a.fallbackDoctorsOnlySubsFilter(ctx)
	}

	logger.Message(ctx, fmt.Sprintf("Количество докторов: %d", doctorsCount))

	return doctorsCount, nil
}

func (a *Action) getDoctorsByCitiesAndSpecialitiesFilter(ctx context.Context, filter dto.Filter) (int64, error) {
	// Получаем всех докторов без лимита, так как у нас стоят индексы и идем сокращать выборку в подписчиков
	orderedIDs, err := a.doctorService.GetDoctorsIDsByFilter(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[ERROR] Ошибка при получении докторов", err)
		return 0, err
	}

	doctorsCount, err := a.subscribersService.FilterDoctorsBySubscribersWithDoctorsIDs(ctx, filter, orderedIDs)
	if err != nil {
		// фолбек
		logger.Error(ctx, "[ERROR] Ошибка при сортировке по подписчиками, делаем фолбек", err)

		return int64(len(orderedIDs)), nil
	}

	return doctorsCount, nil
}

func (a *Action) fallbackDoctorsOnlySubsFilter(ctx context.Context) (int64, error) {
	doctorsCount, err := a.doctorService.GetDoctorsCount(ctx)
	if err != nil {
		return 0, err
	}

	return doctorsCount, nil
}

func (a *Action) defaultFilter(filter dto.Filter) bool {
	if len(filter.Cities) == 0 &&
		len(filter.Specialities) == 0 &&
		len(filter.SocialMedia) == 0 &&
		filter.MaxSubscribers == 5_000_000 &&
		filter.MinSubscribers == 100 {

		return true
	}
	return false
}
