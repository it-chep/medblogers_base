package freelancer

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers/dto"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"strings"
	"sync"

	"github.com/samber/lo"
)

func (s *Service) EnrichFacade(ctx context.Context, dtoMap map[int64]dto.Freelancer, freelancersIDs []int64) {
	// Обогащение специальностями, фотографиями и городами
	var (
		imageMap        map[string]string
		citiesMap       map[int64][]*city.City
		specialitiesMap map[int64][]*speciality.Speciality
		networksMap     map[int64][]*social_network.SocialNetwork
		mu              sync.Mutex
		errs            []error
	)
	g := async.NewGroup()

	// Получаем фотки
	g.Go(func() {
		imgs, err := s.imageEnricher.GetUserPhotos(ctx)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении фотографий: %w", err))
			return
		}
		imageMap = imgs
	})

	// Получаем доп города
	g.Go(func() {
		cities, err := s.additionalStorage.GetAdditionalCities(ctx, freelancersIDs)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении городов: %w", err))
			return
		}
		citiesMap = cities
	})

	// Получаем доп специальности
	g.Go(func() {
		specs, err := s.additionalStorage.GetAdditionalSpecialities(ctx, freelancersIDs)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении специальностей: %w", err))
			return
		}
		specialitiesMap = specs
	})

	// Получаем соц сетей
	g.Go(func() {
		networks, err := s.additionalStorage.GetSocialNetworks(ctx, freelancersIDs)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении специальностей: %w", err))
			return
		}
		networksMap = networks
	})

	g.Wait()

	// Обрабатываем все ошибки
	if len(errs) > 0 {
		for _, e := range errs {
			logger.Error(ctx, "[Filter] Ошибка обогащения", e)
		}
	}

	// обогащаем всеми данными
	enrichImages(ctx, dtoMap, imageMap)
	enrichAdditionalSpecialities(ctx, dtoMap, specialitiesMap)
	enrichAdditionalCities(ctx, dtoMap, citiesMap)
	enrichSocialNetworks(ctx, dtoMap, networksMap)
}

// enrichAdditionalCities - обогащение доп городами
func enrichAdditionalCities(ctx context.Context, freelancersMap map[int64]dto.Freelancer, additionalCitiesMap map[int64][]*city.City) {
	logger.Message(ctx, "[Filter] Обогащение доп городами")

	for freelancerID, cities := range additionalCitiesMap {
		freelancer, ok := freelancersMap[freelancerID]
		if !ok {
			continue
		}

		var builder strings.Builder

		// Сначала ищем основной город
		for _, c := range cities {
			if c.ID() == freelancer.MainCityID {
				builder.WriteString(c.Name())
				break
			}
		}

		counter := 0
		// Добавляем остальные города через запятую
		for _, c := range cities {
			if counter == 2 {
				break
			}
			if int64(c.ID()) != freelancer.MainCityID {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(c.Name())
				counter++
			}
		}

		// Обновляем данные доктора
		if builder.Len() > 0 {
			freelancer.City = builder.String()
			freelancersMap[freelancerID] = freelancer
		}
	}
}

// enrichAdditionalSpecialities - обогащение доп специальностями и доп городами
func enrichAdditionalSpecialities(ctx context.Context, freelancersMap map[int64]dto.Freelancer, additionalSpecialitiesMap map[int64][]*speciality.Speciality) {
	logger.Message(ctx, "[Filter] Обогащение доп специальностями")

	for freelancerID, specialities := range additionalSpecialitiesMap {
		freelancer, ok := freelancersMap[freelancerID]
		if !ok {
			continue
		}

		var builder strings.Builder

		// Сначала ищем основной город
		for _, spec := range specialities {
			if spec.ID() == freelancer.MainSpecialityID {
				builder.WriteString(spec.Name())
				break
			}
		}

		counter := 0
		// Добавляем остальные специальности через запятую
		for _, spec := range specialities {
			if counter == 2 {
				break
			}
			if spec.ID() != freelancer.MainSpecialityID {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(spec.Name())
				counter++
			}
		}

		// Обновляем данные доктора
		if builder.Len() > 0 {
			freelancer.Speciality = builder.String()
			freelancersMap[freelancerID] = freelancer
		}
	}
}

// enrichImages - обогащение фотографиями в миниатюры докторов
func enrichImages(ctx context.Context, freelancersMap map[int64]dto.Freelancer, photos map[string]string) {
	logger.Message(ctx, "[Filter] Обогащение фотографиями")

	for id, freelancer := range freelancersMap {
		photo, ok := photos[freelancer.Slug]
		if !ok {
			// Устанавливаем дефолтное значение
			freelancer.Image = "https://storage.yandexcloud.net/medblogers-photos/zag.jpg"
			freelancersMap[id] = freelancer
			continue
		}

		freelancer.Image = photo
		freelancersMap[id] = freelancer
	}
}

// enrichSocialNetworks - обогащение соц сетями
func enrichSocialNetworks(ctx context.Context, freelancersMap map[int64]dto.Freelancer, socialNetworks map[int64][]*social_network.SocialNetwork) {
	logger.Message(ctx, "[Filter] Обогащение доп соцсетями")

	for freelancerID, networks := range socialNetworks {
		freelancer, ok := freelancersMap[freelancerID]
		if !ok {
			continue
		}

		freelancer.Networks = lo.Map(networks, func(item *social_network.SocialNetwork, _ int) dto.NetworkItem {
			return dto.NetworkItem{
				ID:   item.ID(),
				Name: item.Name(),
				Slug: item.Slug(),
			}
		})

		freelancersMap[freelancerID] = freelancer
	}
}
