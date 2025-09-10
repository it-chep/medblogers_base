package search_freelancers

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/dal"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/dto"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/service/city"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/service/freelancer"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers/service/speciality"
	cityDomain "medblogers_base/internal/modules/freelancers/domain/city"
	specialityDomain "medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	city       *city.Service
	speciality *speciality.Service
	freelancer *freelancer.Service
}

func NewAction(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repository := dal.NewRepository(pool)
	return &Action{
		city:       city.NewSearchService(repository),
		speciality: speciality.NewSearchService(repository),
		freelancer: freelancer.NewSearchService(repository, clients.S3),
	}
}

func (a *Action) Do(ctx context.Context, query string) (dto.SearchDTO, error) {
	logger.Message(ctx, fmt.Sprintf("[Search] Запуск сценария поиска. Query: %s", query))
	searchResult := dto.SearchDTO{}
	// Делаем обычную группу, чтобы если 1 из фильров не работает, он не ломал нам весь сайт
	g := async.NewGroup()

	// получение городов
	g.Go(func() {
		cities, err := a.city.SearchCities(ctx, query)
		if err != nil {
			logger.Error(ctx, "Ошибка при поиске городов", err)
		}
		searchResult.Cities = lo.Map(cities, func(item *cityDomain.City, _ int) dto.CityItem {
			return dto.CityItem{
				ID:               item.ID(),
				Name:             item.Name(),
				FreelancersCount: item.FreelancersCount(),
			}
		})
	})

	// получение специальностей
	g.Go(func() {
		specialities, err := a.speciality.SearchSpecialities(ctx, query)
		if err != nil {
			logger.Error(ctx, "Ошибка при поиске специальностей", err)
		}
		searchResult.Specialities = lo.Map(specialities, func(item *specialityDomain.Speciality, _ int) dto.SpecialityItem {
			return dto.SpecialityItem{
				ID:               item.ID(),
				Name:             item.Name(),
				FreelancersCount: item.FreelancersCount(),
			}
		})
	})

	// получение докторов
	g.Go(func() {
		freelancers, err := a.freelancer.Search(ctx, query)
		if err != nil {
			logger.Error(ctx, "Ошибка при поиске докторов", err)
		}
		searchResult.Freelancers = freelancers
	})

	g.Wait()

	return searchResult, nil
}
