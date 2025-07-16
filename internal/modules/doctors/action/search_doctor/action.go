package search_doctor

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/search_doctor/dal"
	"medblogers_base/internal/modules/doctors/action/search_doctor/dto"
	"medblogers_base/internal/modules/doctors/action/search_doctor/service/city"
	"medblogers_base/internal/modules/doctors/action/search_doctor/service/speciality"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"github.com/samber/lo"

	"medblogers_base/internal/modules/doctors/action/search_doctor/service/doctor"
	cityDomain "medblogers_base/internal/modules/doctors/domain/city"
	specialityDomain "medblogers_base/internal/modules/doctors/domain/speciality"
)

// Action поиск доктора по фио, специальности, города
type Action struct {
	doctor     *doctor.Service
	city       *city.Service
	speciality *speciality.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		doctor:     doctor.NewSearchService(dal.NewRepository(pool), clients.S3),
		city:       city.NewSearchService(dal.NewRepository(pool)),
		speciality: speciality.NewSearchService(dal.NewRepository(pool)),
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
				ID:           int64(item.ID()),
				Name:         item.Name(),
				DoctorsCount: item.DoctorsCount(),
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
				ID:           int64(item.ID()),
				Name:         item.Name(),
				DoctorsCount: item.DoctorsCount(),
			}
		})
	})

	// получение докторов
	g.Go(func() {
		doctors, err := a.doctor.Search(ctx, query)
		if err != nil {
			logger.Error(ctx, "Ошибка при поиске докторов", err)
		}
		searchResult.Doctors = doctors
	})

	g.Wait()

	return searchResult, nil
}
