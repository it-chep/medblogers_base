package validate

import (
	"context"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/service/validate/rules"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/spec"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . CityStorage,SpecialityStorage

type CityStorage interface {
	GetAllCities(ctx context.Context) ([]*city.City, error)
}

type SpecialityStorage interface {
	GetAllSpecialities(ctx context.Context) ([]*speciality.Speciality, error)
}

type Service struct {
	cityStorage       CityStorage
	specialityStorage SpecialityStorage
}

func NewService(cityStorage CityStorage, specialityStorage SpecialityStorage) *Service {
	return &Service{
		cityStorage:       cityStorage,
		specialityStorage: specialityStorage,
	}
}

func (s *Service) ValidateFreelancer(ctx context.Context, createDTO *dto.CreateRequest) ([]dto.ValidationError, error) {
	citiesIDs, err := s.getCitiesIDs(ctx)
	if err != nil {
		return nil, err
	}

	specialitiesIDs, err := s.getSpecialitiesIDs(ctx)
	if err != nil {
		return nil, err
	}

	specification := spec.NewIndependentSpecification[*dto.CreateRequest]().
		And(rules.RuleValidSpecialityID(specialitiesIDs)).
		And(rules.RuleValidSpecialitiesIDs(specialitiesIDs)).
		And(rules.RuleValidCityID(citiesIDs)).
		And(rules.RuleValidAdditionalCitiesIDs(citiesIDs))

	domainErrors := make([]dto.ValidationError, 0)
	for _, validationError := range specification.Validate(ctx, createDTO) {
		var errV dto.ValidationError
		ok := errors.As(validationError, &errV)
		if ok {
			domainErrors = append(domainErrors, dto.ValidationError{
				Field: errV.Field,
				Text:  errV.Text,
			})
			continue
		}
		return nil, err
	}

	return domainErrors, nil
}

func (s *Service) getCitiesIDs(ctx context.Context) ([]int64, error) {
	cities, err := s.cityStorage.GetAllCities(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получении городов при реге", err)
		return nil, err
	}
	citiesIDs := make([]int64, 0, len(cities))
	for _, c := range cities {
		citiesIDs = append(citiesIDs, c.ID())
	}

	return citiesIDs, nil
}

func (s *Service) getSpecialitiesIDs(ctx context.Context) ([]int64, error) {
	specialities, err := s.specialityStorage.GetAllSpecialities(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получении городов при реге", err)
		return nil, err
	}

	specialitiesIDs := make([]int64, 0, len(specialities))
	for _, sp := range specialities {
		specialitiesIDs = append(specialitiesIDs, sp.ID())
	}

	return specialitiesIDs, nil
}
