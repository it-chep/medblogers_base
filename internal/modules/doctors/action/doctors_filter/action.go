package doctors_filter

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dal"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/doctor"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/page"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/subscribers"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action фильтрация докторов
type Action struct {
	subscribersFilter *subscribers.Service
	doctorsFilter     *doctor.Service
	pageService       *page.Service
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
		pageService: page.New(),
	}
}

// Do - фильтрует докторов
// Если есть город или специальность в фильтре, то нам надо получить сначала данные из базы по городу и специальности
// По этим ID сходить в подписчики и отфильтровать там подписчиков, тк если мы лимитом получим 30 из подписчиков,
// то мы можем не добрать позже при фильтрации по городам и специальностям
func (a Action) Do(ctx context.Context, filter dto.Filter) (dto.Response, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[Filter] Фильтрация докторов: MaxSubs: %d, MinSubs: %d, SocialMedia: %v, Cities: %v, Specialities: %v, Page: %d, Sort: %s",
		filter.MinSubscribers, filter.MinSubscribers, filter.SocialMedia, filter.Cities, filter.Specialities, filter.Page, filter.Sort,
	))

	// Получение по фильтрам из базы если есть
	if len(filter.Cities) != 0 || len(filter.Specialities) != 0 {
		// Получаем всех докторов без лимита, так как у нас стоят индексы и идем сокращать выборку в подписчиков
		doctorsMap, err := a.doctorsFilter.GetDoctorsByFilter(ctx, filter)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при получении докторов", err)
			return dto.Response{}, err
		}

		// Фильтруем и сортируем в подписчиках
		subsResponse, err := a.subscribersFilter.FilterDoctorsBySubscribersWithDoctorsIDs(ctx, filter, lo.Keys(doctorsMap))
		if err != nil {
			return dto.Response{}, err
		}

		// Обогащаем данными только нужных докторов
		a.doctorsFilter.EnrichFacade(ctx, doctorsMap, lo.Keys(subsResponse.Doctors))

		// Маппим данные подписчиков и докторов
		mappedDoctors := a.subscribersFilter.MapDoctorsWithSubscribers(doctorsMap, subsResponse.Doctors, subsResponse.OrderedIDs)
		pagesCount := a.pageService.GetPagesCountBySubscribersFilter(subsResponse.DoctorsCount)

		return dto.Response{
			Doctors:          mappedDoctors,
			CurrentPage:      filter.Page,
			Pages:            pagesCount,
			SubscribersCount: subsResponse.SubsCount,
		}, nil
	}

	// Фильтр только по подписчикам
	logger.Message(ctx, "[Filter] Фильтруем по подписчикам")
	subsResponse, err := a.subscribersFilter.FilterDoctorsBySubscribers(ctx, filter)
	if err != nil {
		logger.Error(ctx, "[Filter] Ошибка при фильтрации докторов по подписчиками, делаем дефолт", err)
	}

	doctorsMap, err := a.doctorsFilter.GetDoctorsByIDs(ctx, filter.Page, subsResponse.OrderedIDs)
	if err != nil {
		logger.Error(ctx, "[Filter] Ошибка при обогащении данными", err)
	}

	// Обогащаем данными только нужных докторов
	a.doctorsFilter.EnrichFacade(ctx, doctorsMap, lo.Keys(subsResponse.Doctors))

	// Маппим данные подписчиков и докторов
	mappedDoctors := a.subscribersFilter.MapDoctorsWithSubscribers(doctorsMap, subsResponse.Doctors, subsResponse.OrderedIDs)
	pagesCount := a.pageService.GetPagesCountBySubscribersFilter(subsResponse.DoctorsCount) // todo учесть кейс с вкл/выкл врача

	return dto.Response{
		Doctors:          mappedDoctors,
		CurrentPage:      filter.Page,
		Pages:            pagesCount,
		SubscribersCount: subsResponse.SubsCount,
	}, nil
}
