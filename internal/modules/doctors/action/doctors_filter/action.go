package doctors_filter

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dal"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/doctor"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/subscribers"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/vip"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/samber/lo"
)

// todo как-то улучшить а то выглядит говном

// Action фильтрация докторов
type Action struct {
	subscribersFilter *subscribers.Service
	doctorsFilter     *doctor.Service
	vipService        *vip.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repo := dal.NewRepository(pool)
	return &Action{
		subscribersFilter: subscribers.New(clients.Subscribers),
		doctorsFilter:     doctor.New(repo, repo, clients.S3, clients.Subscribers),
		vipService:        vip.NewService(repo),
	}
}

// Do - фильтрует докторов
// Если есть город или специальность в фильтре, то нам надо получить сначала данные из базы по городу и специальности
// По этим ID сходить в подписчики и отфильтровать там подписчиков, тк если мы лимитом получим 30 из подписчиков,
// то мы можем не добрать позже при фильтрации по городам и специальностям
func (a Action) Do(ctx context.Context, filter dto.Filter) (dto.Response, error) {
	if len(filter.Cities) != 0 || len(filter.Specialities) != 0 {
		return a.getDoctorsByCitiesAndSpecialitiesFilter(ctx, filter)
	}

	subsResponse, err := a.subscribersFilter.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[ERROR] Ошибка при фильтрации докторов по подписчиками, делаем фолбек", err)
		return a.fallbackDoctorsOnlySubsFilter(ctx, filter)
	}

	doctorsMap, err := a.doctorsFilter.GetDoctorsByIDs(ctx, filter.Page, subsResponse.OrderedIDs)
	if err != nil {
		logger.Error(ctx, "[ERROR] Ошибка при обогащении данными", err)
	}

	// Обогащаем данными только нужных докторов
	a.doctorsFilter.EnrichFacade(ctx, doctorsMap, lo.Keys(subsResponse.Doctors))
	mappedDoctors := a.subscribersFilter.MapDoctorsWithSubscribers(doctorsMap, subsResponse.Doctors, subsResponse.OrderedIDs)

	return dto.Response{
		Doctors:          mappedDoctors,
		Vip:              a.vipService.GetDoctorsVipInfo(ctx, mappedDoctors.GetVipIDs()),
		SubscribersCount: subsResponse.SubsCount,
	}, nil
}

// Получаем всех докторов без лимита, так как у нас стоят индексы и идем сокращать выборку в подписчиков
func (a Action) getDoctorsByCitiesAndSpecialitiesFilter(ctx context.Context, filter dto.Filter) (dto.Response, error) {
	doctorsMap, orderedIDs, err := a.doctorsFilter.GetDoctorsByFilter(ctx, filter)
	if err != nil {
		return dto.Response{}, err
	}

	subsResponse, err := a.subscribersFilter.FilterDoctorsBySubscribersWithDoctorsIDs(ctx, filter, orderedIDs)
	if err != nil {
		a.doctorsFilter.EnrichFacade(ctx, doctorsMap, orderedIDs)
		trimmedDoctors := a.doctorsFilter.TrimFallbackDoctors(filter, doctorsMap, orderedIDs)
		return dto.Response{
			Doctors: trimmedDoctors,
			Vip:     a.vipService.GetDoctorsVipInfo(ctx, trimmedDoctors.GetVipIDs()),
		}, nil
	}

	// Обогащаем данными только нужных докторов
	a.doctorsFilter.EnrichFacade(ctx, doctorsMap, lo.Keys(subsResponse.Doctors))
	mappedDoctors := a.subscribersFilter.MapDoctorsWithSubscribers(doctorsMap, subsResponse.Doctors, subsResponse.OrderedIDs)

	return dto.Response{
		Doctors:          mappedDoctors,
		Vip:              a.vipService.GetDoctorsVipInfo(ctx, mappedDoctors.GetVipIDs()),
		SubscribersCount: subsResponse.SubsCount,
	}, nil
}

func (a Action) fallbackDoctorsOnlySubsFilter(ctx context.Context, filter dto.Filter) (dto.Response, error) {
	doctors, err := a.doctorsFilter.GetDoctors(ctx, filter.Page)
	if err != nil {
		return dto.Response{}, err
	}

	return dto.Response{
		Doctors: doctors,
		Vip:     a.vipService.GetDoctorsVipInfo(ctx, doctors.GetVipIDs()),
	}, nil
}
