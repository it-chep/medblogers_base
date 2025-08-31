package doctor

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"strings"
	"sync"

	"github.com/samber/lo"
)

func (s *Service) EnrichFacade(ctx context.Context, dtoMap map[int64]dto.Doctor, doctorsIDs []int64) {
	// Обогащение специальностями, фотографиями и городами
	var (
		imageMap        map[string]string
		citiesMap       map[int64][]*city.City
		specialitiesMap map[int64][]*speciality.Speciality
		mu              sync.Mutex
		errs            []error
	)
	g := async.NewGroup()

	// Получаем фотки
	g.Go(func() {
		imgs, err := s.imageGetter.GetUserPhotos(ctx)
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
		cities, err := s.additionalStorage.GetDoctorAdditionalCities(ctx, doctorsIDs)
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
		specs, err := s.additionalStorage.GetDoctorAdditionalSpecialities(ctx, doctorsIDs)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении специальностей: %w", err))
			return
		}
		specialitiesMap = specs
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
}

// enrichSubscribers - обогащение подписчиками в миниатюры докторов
func enrichSubscribers(ctx context.Context, doctorsMap map[int64]dto.Doctor, subscribersMap map[int64]indto.GetSubscribersByDoctorIDsResponse) {
	logger.Message(ctx, "[Filter] Обогащение подписчиками")

	for id, doc := range doctorsMap {
		subsInfo, ok := subscribersMap[doc.ID]
		if !ok {
			continue
		}

		doc.TgSubsCount = subsInfo.TgSubsCount
		doc.TgSubsCountText = subsInfo.TgSubsCountText
		doc.InstSubsCount = subsInfo.InstSubsCount
		doc.InstSubsCountText = subsInfo.InstSubsCountText

		doctorsMap[id] = doc
	}
}

// enrichAdditionalCities - обогащение доп городами
func enrichAdditionalCities(ctx context.Context, doctorsMap map[int64]dto.Doctor, additionalCitiesMap map[int64][]*city.City) {
	logger.Message(ctx, "[Filter] Обогащение доп городами")

	for doctorID, cities := range additionalCitiesMap {
		doctor, ok := doctorsMap[doctorID]
		if !ok {
			continue
		}

		var builder strings.Builder

		// Сначала ищем основной город
		for _, c := range cities {
			if int64(c.ID()) == doctor.MainCityID {
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
			if int64(c.ID()) != doctor.MainCityID {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(c.Name())
				counter++
			}
		}

		// Обновляем данные доктора
		if builder.Len() > 0 {
			doctor.City = builder.String()
			doctorsMap[doctorID] = doctor
		}
	}
}

// enrichAdditionalSpecialities - обогащение доп специальностями и доп городами
func enrichAdditionalSpecialities(ctx context.Context, doctorsMap map[int64]dto.Doctor, additionalSpecialitiesMap map[int64][]*speciality.Speciality) {
	logger.Message(ctx, "[Filter] Обогащение доп специальностями")

	for doctorID, specialities := range additionalSpecialitiesMap {
		doctor, ok := doctorsMap[doctorID]
		if !ok {
			continue
		}

		var builder strings.Builder

		// Сначала ищем основной город
		for _, spec := range specialities {
			if int64(spec.ID()) == doctor.MainSpecialityID {
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
			if int64(spec.ID()) != doctor.MainSpecialityID {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(spec.Name())
				counter++
			}
		}

		// Обновляем данные доктора
		if builder.Len() > 0 {
			doctor.Speciality = builder.String()
			doctorsMap[doctorID] = doctor
		}
	}
}

// enrichImages - обогащение фотографиями в миниатюры докторов
func enrichImages(ctx context.Context, doctorsMap map[int64]dto.Doctor, photos map[string]string) {
	logger.Message(ctx, "[Filter] Обогащение фотографиями")

	for id, doc := range doctorsMap {
		photo, ok := photos[doc.Slug]
		if !ok {
			// Устанавливаем дефолтное значение
			doc.Image = "https://storage.yandexcloud.net/medblogers-photos/zag.jpg"
			doctorsMap[id] = doc
			continue
		}

		doc.Image = photo
		doctorsMap[id] = doc
	}
}

func (s *Service) enrichWithSubscribersFacade(ctx context.Context, doctorsMap map[int64]dto.Doctor) {
	var (
		subscribersMap  map[int64]indto.GetSubscribersByDoctorIDsResponse
		imageMap        map[string]string
		citiesMap       map[int64][]*city.City
		specialitiesMap map[int64][]*speciality.Speciality
		mu              sync.Mutex
		errs            []error
	)
	g := async.NewGroup()

	// Получаем количество подписчиков
	g.Go(func() {
		subs, err := s.subscribersGetter.GetSubscribersByDoctorIDs(ctx, lo.Keys(doctorsMap))
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении подписчиков: %w", err))
			return
		}
		subscribersMap = subs
	})

	// Получаем фотки
	g.Go(func() {
		imgs, err := s.imageGetter.GetUserPhotos(ctx)
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
		cities, err := s.additionalStorage.GetDoctorAdditionalCities(ctx, lo.Keys(doctorsMap))
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
		specs, err := s.additionalStorage.GetDoctorAdditionalSpecialities(ctx, lo.Keys(doctorsMap))
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при получении специальностей: %w", err))
			return
		}
		specialitiesMap = specs
	})

	g.Wait()

	// Обрабатываем все ошибки
	if len(errs) > 0 {
		for _, e := range errs {
			logger.Error(ctx, "[Filter] Ошибка обогащения", e)
		}
	}

	// обогащаем всеми данными
	enrichImages(ctx, doctorsMap, imageMap)
	enrichAdditionalSpecialities(ctx, doctorsMap, specialitiesMap)
	enrichAdditionalCities(ctx, doctorsMap, citiesMap)
	enrichSubscribers(ctx, doctorsMap, subscribersMap)
}
