package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers/dto"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
)

type Storage interface {
	FilterFreelancers(ctx context.Context, filter freelancer.Filter) (map[int64]*freelancer.Freelancer, []int64, error)
	GetFreelancers(ctx context.Context, filter freelancer.Filter) (map[int64]*freelancer.Freelancer, []int64, error)
}

type ImageEnricher interface {
	GetUserPhotos(ctx context.Context) (map[string]string, error)
}

type AdditionalStorage interface {
	GetAdditionalCities(ctx context.Context, medblogersIDs []int64) (map[int64][]*city.City, error)
	GetAdditionalSpecialities(ctx context.Context, medblogersIDs []int64) (map[int64][]*speciality.Speciality, error)
	GetSocialNetworks(ctx context.Context, medblogersIDs []int64) (map[int64][]*social_network.SocialNetwork, error)
}

type Service struct {
	storage           Storage
	imageEnricher     ImageEnricher
	additionalStorage AdditionalStorage
}

func NewService(storage Storage, imageEnricher ImageEnricher, additionalStorage AdditionalStorage) *Service {
	return &Service{
		storage:           storage,
		imageEnricher:     imageEnricher,
		additionalStorage: additionalStorage,
	}
}

// GetFreelancersByFilter - фильтрация фрилансеров по полям в базе
func (s *Service) GetFreelancersByFilter(ctx context.Context, filter freelancer.Filter) ([]dto.Freelancer, error) {
	logger.Message(ctx, "[Filter][Service] Получение фрилансеров по фильтрам")
	freelancersMap, orderedIDs, err := s.storage.FilterFreelancers(ctx, filter)
	if err != nil {
		return nil, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(freelancersMap)

	// обогащение необходимыми сущностями
	s.EnrichFacade(ctx, dtoMap, orderedIDs)

	// Делаем правильный порядок докторов
	result := make([]dto.Freelancer, 0, len(dtoMap))
	// Итерируемся по map напрямую
	for _, id := range orderedIDs {
		result = append(result, dtoMap[id])
	}

	return result, nil
}

// GetFreelancers - дефолтное получение фрилансеров без фильтров
func (s *Service) GetFreelancers(ctx context.Context, filter freelancer.Filter) ([]dto.Freelancer, error) {
	logger.Message(ctx, "[Filter][Service] Дефолтное получение фрилансеров")
	// Получаем докторов
	freelancersMap, orderedIDs, err := s.storage.GetFreelancers(ctx, filter)
	if err != nil {
		return nil, err
	}

	// конвертация в DTO
	dtoMap := s.convertToDTOMap(freelancersMap)

	// обогащение необходимыми сущностями
	s.EnrichFacade(ctx, dtoMap, orderedIDs)

	// Делаем правильный порядок докторов
	result := make([]dto.Freelancer, 0, len(dtoMap))
	// Итерируемся по map напрямую
	for _, id := range orderedIDs {
		result = append(result, dtoMap[id])
	}

	return result, nil
}

func (s *Service) convertToDTOMap(freelancersMap map[int64]*freelancer.Freelancer) map[int64]dto.Freelancer {
	dtoMap := make(map[int64]dto.Freelancer, len(freelancersMap))
	for _, freelanc := range freelancersMap {
		dtoMap[freelanc.GetID()] = dto.Freelancer{
			ID:                   freelanc.GetID(),
			Slug:                 freelanc.GetSlug(),
			Name:                 freelanc.GetName(),
			MainCityID:           freelanc.GetMainCityID(),
			MainSpecialityID:     freelanc.GetMainSpecialityID(),
			PriceCategory:        freelanc.GetPriceCategory(),
			AgencyRepresentative: freelanc.IsAgencyRepresentative(),
		}
	}

	return dtoMap
}
