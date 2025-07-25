package doctors_filter

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dal"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/doctor"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/page"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/service/subscribers"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/samber/lo"
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
		pageService: page.New(dal.NewRepository(pool)),
	}
}

func (a Action) Do(ctx context.Context, filter dto.Filter) (dto.Response, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[Filter] Фильтрация докторов: MaxSubs: %d, MinSubs: %d, SocialMedia: %v, Cities: %v, Specialities: %v, Page: %d",
		filter.MinSubscribers, filter.MinSubscribers, filter.SocialMedia, filter.Cities, filter.Specialities, filter.Page,
	))

	//	Если фильтр не подписчиков. Берем докторов, обогажаем миниатюрами
	if filter.MinSubscribers == 0 && filter.MaxSubscribers == 0 && len(filter.SocialMedia) == 0 {
		doctorsMap, err := a.doctorsFilter.GetDoctorsByFilter(ctx, filter)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при получении докторов", err)
			return dto.Response{}, err
		}

		pagesCount, err := a.pageService.GetPagesCount(ctx, filter)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при подсчете количества страниц", err)
		}
		return dto.Response{
			Doctors:     lo.Values(doctorsMap),
			CurrentPage: filter.Page,
			Pages:       pagesCount,
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
	pagesCount := a.pageService.GetPagesCountBySubscribersFilter(int64(len(lo.Keys(subsResponse))))

	return dto.Response{
		Doctors:     lo.Values(doctorsMap),
		CurrentPage: filter.Page,
		Pages:       pagesCount,
	}, nil
}
