package doctor

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"strings"
)

// enrichSubscribers - обогащение подписчиками в миниатюры докторов
func (s *Service) enrichSubscribers(ctx context.Context, doctorsMap map[int64]dto.Doctor, subscribersMap map[int64]indto.GetSubscribersByDoctorIDsResponse) {
	logger.Message(ctx, "[Filter] Обогащение подписчиками")

	for _, doc := range lo.Values(doctorsMap) {
		subsInfo, ok := subscribersMap[doc.ID]
		if !ok {
			continue
		}

		doc.TgSubsCount = subsInfo.TgSubsCount
		doc.TgSubsCountText = subsInfo.TgSubsCountText
		doc.InstSubsCount = subsInfo.InstSubsCount
		doc.InstSubsCountText = subsInfo.InstSubsCountText
	}
}

// enrichAdditionalCities - обогащение доп городами
func (s *Service) enrichAdditionalCities(ctx context.Context, doctorsMap map[int64]dto.Doctor, additionalCitiesMap map[int64][]*city.City) {
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

		// Добавляем остальные города через запятую
		for _, c := range cities {
			if int64(c.ID()) != doctor.MainCityID {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(c.Name())
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
func (s *Service) enrichAdditionalSpecialities(ctx context.Context, doctorsMap map[int64]dto.Doctor, additionalSpecialitiesMap map[int64][]*speciality.Speciality) {
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

		// Добавляем остальные специальности через запятую
		for _, spec := range specialities {
			if int64(spec.ID()) != doctor.MainSpecialityID {
				if builder.Len() > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(spec.Name())
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
func (s *Service) enrichImages(ctx context.Context, doctorsMap map[int64]dto.Doctor, photos map[string]string) {
	logger.Message(ctx, "[Filter] Обогащение фотографиями")

	for _, doc := range lo.Values(doctorsMap) {
		photo, ok := photos[doc.Slug]
		if !ok {
			// todo дефолт значение
			continue
		}

		doc.Image = photo
	}
}

func (s *Service) enrichFacade(ctx context.Context, doctorsMap map[int64]dto.Doctor) {
	var (
		err             error
		subscribersMap  = make(map[int64]indto.GetSubscribersByDoctorIDsResponse)
		imageMap        = make(map[string]string)
		citiesMap       = make(map[int64][]*city.City)
		specialitiesMap = make(map[int64][]*speciality.Speciality)
	)
	g := async.NewGroup()

	// Получаем количество подписчиков
	g.Go(func() {
		subscribersMap, err = s.subscribersGetter.GetSubscribersByDoctorIDs(ctx, lo.Keys(doctorsMap))
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при обогащении подписчиками", err)
		}
	})

	// Получаем фотки
	g.Go(func() {
		imageMap, err = s.imageGetter.GetUserPhotos(ctx)
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при обогащении фотографией", err)
		}
	})

	// Получаем доп города
	g.Go(func() {
		citiesMap, err = s.additionalStorage.GetDoctorAdditionalCities(ctx, lo.Keys(doctorsMap))
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при обогащении доп специальностями и городами", err)
		}
	})

	// Получаем доп специальности
	g.Go(func() {
		specialitiesMap, err = s.additionalStorage.GetDoctorAdditionalSpecialities(ctx, lo.Keys(doctorsMap))
		if err != nil {
			logger.Error(ctx, "[Filter] Ошибка при обогащении доп специальностями и городами", err)
		}
	})

	g.Wait()

	// обогащаем всеми данными
	s.enrichImages(ctx, doctorsMap, imageMap)
	s.enrichAdditionalSpecialities(ctx, doctorsMap, specialitiesMap)
	s.enrichAdditionalCities(ctx, doctorsMap, citiesMap)
	s.enrichSubscribers(ctx, doctorsMap, subscribersMap)
}
