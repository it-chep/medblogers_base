package doctor

import (
	"context"
	consts "medblogers_base/internal/dto"
	"medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	"medblogers_base/internal/modules/doctors/client/subscribers/indto"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage,ImageEnricher,SubscribersEnricher,AdditionalStorage

type Storage interface {
	FilterDoctors(ctx context.Context, filter dto.Filter) (map[doctor.MedblogersID]*doctor.Doctor, error)
	GetDoctors(ctx context.Context, currentPage int64) (map[doctor.MedblogersID]*doctor.Doctor, error)
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

type CommonDal interface {
	GetDoctorsCount(ctx context.Context) (int64, error)
}

type Service struct {
	storage           Storage
	additionalStorage AdditionalStorage
	imageGetter       ImageEnricher
	subscribersGetter SubscribersEnricher
	commonDal         CommonDal
}

func New(storage Storage, additionalStorage AdditionalStorage, imageGetter ImageEnricher, subscribersGetter SubscribersEnricher, commonDal CommonDal) *Service {
	return &Service{
		storage:           storage,
		additionalStorage: additionalStorage,
		imageGetter:       imageGetter,
		subscribersGetter: subscribersGetter,
		commonDal:         commonDal,
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

	return dtoMap, nil
}

// GetDoctors - дефолтное получение докторов без фильтров
func (s *Service) GetDoctors(ctx context.Context, currentPage int64) (map[int64]dto.Doctor, int64, error) {
	logger.Message(ctx, "[Filter][Service] Дефолтное получение докторов")
	// Получаем докторов
	doctorsMap, err := s.storage.GetDoctors(ctx, currentPage)
	if err != nil {
		return nil, 0, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(doctorsMap)

	doctorsCount, err := s.commonDal.GetDoctorsCount(ctx)
	if err != nil {
		return dtoMap, 0, err
	}

	// обогащение необходимыми сущностями
	s.enrichWithSubscribersFacade(ctx, dtoMap)

	return dtoMap, doctorsCount, nil
}

// GetDoctorsByIDs - получение докторов по переданным IDs
func (s *Service) GetDoctorsByIDs(ctx context.Context, currentPage int64, ids []int64) (map[int64]dto.Doctor, error) {
	logger.Message(ctx, "[Filter][Service] Получение докторов по IDs")

	// Запрос докторов
	doctorsMap, err := s.storage.GetDoctorsByIDs(ctx, currentPage, ids)
	if err != nil {
		return nil, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(doctorsMap)

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

func (s *Service) TrimFallbackDoctors(filter dto.Filter, doctorsMap map[int64]dto.Doctor) map[int64]dto.Doctor {
	if len(doctorsMap) == 0 {
		return doctorsMap
	}

	offset := (filter.Page - 1) * consts.LimitDoctorsOnPage
	limit := consts.LimitDoctorsOnPage

	if offset >= int64(len(doctorsMap)) {
		return doctorsMap
	}

	if offset == 0 && limit >= int64(len(doctorsMap)) {
		return doctorsMap
	}

	result := make(map[int64]dto.Doctor, limit)
	count := int64(0)

	// Итерируемся по map напрямую
	for id, doc := range doctorsMap {
		if count >= offset && count < offset+limit {
			result[id] = doc
		}
		count++

		// Прерываем если набрали нужное количество
		if int64(len(result)) >= limit {
			break
		}
	}

	return result
}
