package doctors

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/action/get_recommendations/dto"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/doctor"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"strings"
	"sync"
)

func (s *Service) EnrichFacade(ctx context.Context, dtoMap map[int64]dto.Doctor, doctorsIDs []int64) {
	// Обогащение специальностями, фотографиями и городами
	var (
		imageMap        map[doctor.S3Key]string
		citiesMap       map[int64][]*city.City
		specialitiesMap map[int64][]*speciality.Speciality
		mu              sync.Mutex
		errs            []error
	)
	g := async.NewGroup()

	// Получаем фотки
	g.Go(func() {
		imgs, err := s.imageGetter.GetDoctorsPhotos(ctx)
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
		cities, err := s.dal.GetDoctorAdditionalCities(ctx, doctorsIDs)
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
		specs, err := s.dal.GetDoctorAdditionalSpecialities(ctx, doctorsIDs)
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
	enrichImages(dtoMap, imageMap)
	enrichAdditionalSpecialities(dtoMap, specialitiesMap)
	enrichAdditionalCities(dtoMap, citiesMap)
}

// enrichAdditionalCities - обогащение доп городами
func enrichAdditionalCities(doctorsMap map[int64]dto.Doctor, additionalCitiesMap map[int64][]*city.City) {

	for doctorID, cities := range additionalCitiesMap {
		doctor, ok := doctorsMap[doctorID]
		if !ok {
			continue
		}

		var builder strings.Builder

		// Сначала ищем основной город
		for _, c := range cities {
			if c.ID() == doctor.MainCityID {
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
			if c.ID() != doctor.MainCityID {
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
func enrichAdditionalSpecialities(doctorsMap map[int64]dto.Doctor, additionalSpecialitiesMap map[int64][]*speciality.Speciality) {
	for doctorID, specialities := range additionalSpecialitiesMap {
		doctor, ok := doctorsMap[doctorID]
		if !ok {
			continue
		}

		var builder strings.Builder

		// Сначала ищем основной город
		for _, spec := range specialities {
			if spec.ID() == doctor.MainSpecialityID {
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
			if spec.ID() != doctor.MainSpecialityID {
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
func enrichImages(doctorsMap map[int64]dto.Doctor, photos map[doctor.S3Key]string) {
	for id, doc := range doctorsMap {
		photo, ok := photos[doctor.S3Key(doc.S3Key)]
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
