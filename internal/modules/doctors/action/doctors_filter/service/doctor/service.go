package doctor

import (
	"context"
	"fmt"
	consts "medblogers_base/internal/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"sync"

	"github.com/samber/lo"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage,ImageEnricher,SubscribersEnricher,AdditionalStorage

type Storage interface {
	FilterDoctors(ctx context.Context, filter dto.Filter) (map[doctor.MedblogersID]*doctor.Doctor, error)
	GetDoctors(ctx context.Context, limit, offset int64) (map[doctor.MedblogersID]*doctor.Doctor, error)
	GetDoctorsByIDs(ctx context.Context, currentPage int64, ids []int64) (map[doctor.MedblogersID]*doctor.Doctor, error)
}

type ImageEnricher interface {
	GetUserPhotos(ctx context.Context) (map[string]string, error)
}

type SubscribersEnricher interface {
	GetSubscribersByDoctorIDs(ctx context.Context, medblogersIDs []int64) (map[int64]indto.GetSubscribersByDoctorIDsResponse, error)
}

type AdditionalStorage interface {
	GetDoctorAdditionalCities(ctx context.Context, medblogersIDs []int64) (map[int64][]*city.City, error)
	GetDoctorAdditionalSpecialities(ctx context.Context, medblogersIDs []int64) (map[int64][]*speciality.Speciality, error)
}

type Service struct {
	storage           Storage
	additionalStorage AdditionalStorage
	imageGetter       ImageEnricher
	subscribersGetter SubscribersEnricher
}

func New(storage Storage, additionalStorage AdditionalStorage, imageGetter ImageEnricher, subscribersGetter SubscribersEnricher) *Service {
	return &Service{
		storage:           storage,
		additionalStorage: additionalStorage,
		imageGetter:       imageGetter,
		subscribersGetter: subscribersGetter,
	}
}

// GetDoctorsByFilter - фильтрация докторов по полям в базе
func (s *Service) GetDoctorsByFilter(ctx context.Context, filter dto.Filter) (map[int64]dto.Doctor, error) {
	logger.Message(ctx, "[Filter][Service] Получение докторов по фильтрам")
	doctorsMap, err := s.storage.FilterDoctors(ctx, filter)
	if err != nil {
		return nil, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(doctorsMap)

	// обогащение необходимыми сущностями
	s.enrichFacade(ctx, dtoMap)

	return dtoMap, nil
}

// GetDoctors - дефолтное получение докторов без фильтров
func (s *Service) GetDoctors(ctx context.Context, currentPage int64) (map[int64]dto.Doctor, error) {
	logger.Message(ctx, "[Filter][Service] Дефолтное получение докторов")
	// Получаем докторов
	doctorsMap, err := s.storage.GetDoctors(ctx, consts.LimitDoctorsOnPage, currentPage*consts.LimitDoctorsOnPage)
	if err != nil {
		return nil, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(doctorsMap)

	// обогащение необходимыми сущностями
	s.enrichFacade(ctx, dtoMap)

	return dtoMap, nil
}

// GetDoctorsByIDs - получение докторов по переданным IDs
func (s *Service) GetDoctorsByIDs(ctx context.Context, currentPage int64, ids []int64) (map[int64]dto.Doctor, error) {
	if len(ids) == 0 {
		return s.GetDoctors(ctx, currentPage)
	}
	logger.Message(ctx, "[Filter][Service] Получение докторов по IDs")

	// Запрос докторов
	doctorsMap, err := s.storage.GetDoctorsByIDs(ctx, currentPage, ids)
	if err != nil {
		return nil, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(doctorsMap)
	// todo переделать, пиздец не нравится
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
		cities, err := s.additionalStorage.GetDoctorAdditionalCities(ctx, lo.Keys(dtoMap))
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
		specs, err := s.additionalStorage.GetDoctorAdditionalSpecialities(ctx, lo.Keys(dtoMap))
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
	s.enrichImages(ctx, dtoMap, imageMap)
	s.enrichAdditionalSpecialities(ctx, dtoMap, specialitiesMap)
	s.enrichAdditionalCities(ctx, dtoMap, citiesMap)

	return dtoMap, nil
}

func (s *Service) convertToDTOMap(doctorsMap map[doctor.MedblogersID]*doctor.Doctor) map[int64]dto.Doctor {
	dtoMap := make(map[int64]dto.Doctor, len(doctorsMap))
	for _, doc := range doctorsMap {
		dtoMap[int64(doc.GetID())] = dto.Doctor{
			ID:               int64(doc.GetID()),
			Slug:             doc.GetSlug(),
			Name:             doc.GetName(),
			InstLink:         doc.GetInstURL(),
			TgLink:           doc.GetTgChannelURL(),
			MainCityID:       int64(doc.GetMainCityID()),
			MainSpecialityID: int64(doc.GetMainSpecialityID()),
		}
	}

	return dtoMap
}
