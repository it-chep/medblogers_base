package validate

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/validate/rules"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
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

func (s *Service) ValidateDoctor(ctx context.Context, createDTO dto.CreateDoctorRequest) ([]dto.ValidationError, error) {
	citiesIDs, err := s.getCitiesIDs(ctx)
	if err != nil {
		return nil, err
	}

	specialitiesIDs, err := s.getSpecialitiesIDs(ctx)
	if err != nil {
		return nil, err
	}

	specification := spec.NewIndependentSpecification[*dto.CreateDoctorRequest]()

	specification.And(rules.RuleValidSpecialityID(specialitiesIDs))
	specification.And(rules.RuleValidSpecialitiesIDs(specialitiesIDs))
	specification.And(rules.RuleValidCityID(citiesIDs))
	specification.And(rules.RuleValidAdditionalCitiesIDs(citiesIDs))
	specification.And(rules.RuleAtLeastOneSocialMedia())
	specification.And(rules.RuleValidSiteLink())
	specification.And(rules.RuleValidBirthDate())

	specification.Validate(ctx, &createDTO)
	validationErrors := specification.Errors()
	if len(validationErrors) > 0 {
		domainErrors := make([]dto.ValidationError, 0, len(validationErrors))
		for _, validationError := range validationErrors {
			domainErrors = append(domainErrors, dto.ValidationError{
				Field: validationError.Field,
				Text:  validationError.Message,
			})
		}

		return domainErrors, nil
	}

	return []dto.ValidationError{}, nil
}

func (s *Service) getCitiesIDs(ctx context.Context) ([]int64, error) {
	cities, err := s.cityStorage.GetAllCities(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка получении городов при реге", err)
		return nil, err
	}
	citiesIDs := make([]int64, 0, len(cities))
	for _, c := range cities {
		citiesIDs = append(citiesIDs, int64(c.ID()))
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
		specialitiesIDs = append(specialitiesIDs, int64(sp.ID()))
	}

	return specialitiesIDs, nil
}
