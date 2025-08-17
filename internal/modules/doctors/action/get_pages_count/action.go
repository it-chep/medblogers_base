package get_pages_count

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/dal"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/dto"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/service/doctors"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/service/page"
	"medblogers_base/internal/modules/doctors/action/get_pages_count/service/subscribers"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	pageService        *page.Service
	doctorsService     *doctors.Service
	subscribersService *subscribers.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		subscribersService: subscribers.New(clients.Subscribers),
		doctorsService:     doctors.New(dal.NewRepository(pool), doctor_dal.NewRepository(pool)),
		pageService:        page.New(),
	}
}

func (a *Action) Do(ctx context.Context, filter dto.Filter) (int64, error) {
	// Получение по фильтрам из базы если есть
	if len(filter.Cities) != 0 || len(filter.Specialities) != 0 {
		return a.getDoctorsByCitiesAndSpecialitiesFilter(ctx, filter)
	}

	// Фильтр только по подписчикам
	subsDoctorsCount, err := a.subscribersService.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[ERROR] Ошибка при фильтрации докторов по подписчиками, делаем фолбек", err)
		return a.fallbackDoctorsOnlySubsFilter(ctx)
	}

	return a.pageService.GetPagesCount(subsDoctorsCount), nil
}

func (a *Action) getDoctorsByCitiesAndSpecialitiesFilter(ctx context.Context, filter dto.Filter) (int64, error) {
	// Получаем всех докторов без лимита, так как у нас стоят индексы и идем сокращать выборку в подписчиков
	orderedIDs, err := a.doctorsService.GetDoctorsIDsByFilter(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[ERROR] Ошибка при получении докторов", err)
		return 0, err
	}

	subsDoctorsCount, err := a.subscribersService.FilterDoctorsBySubscribersWithDoctorsIDs(ctx, filter, orderedIDs)
	if err != nil {
		// фолбек
		logger.Error(ctx, "[ERROR] Ошибка при сортировке по подписчиками, делаем фолбек", err)

		return a.pageService.GetPagesCount(int64(len(orderedIDs))), nil
	}

	return a.pageService.GetPagesCount(subsDoctorsCount), nil
}

func (a *Action) fallbackDoctorsOnlySubsFilter(ctx context.Context) (int64, error) {
	doctorsCount, err := a.doctorsService.GetDoctorsCount(ctx)
	if err != nil {
		return 0, err
	}

	return a.pageService.GetPagesCount(doctorsCount), nil
}
