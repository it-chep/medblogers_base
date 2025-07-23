package validate

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/validate/rules"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage,CityStorage,SpecialityStorage

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

func (s *Service) ValidateDoctor(ctx context.Context, createDTO dto.CreateDoctorRequest) (_ dto.CreateDoctorRequest, err error) {
	citiesIDs, err := s.getCitiesIDs(ctx)
	if err != nil {
		return createDTO, err
	}

	specialitiesIDs, err := s.getSpecialitiesIDs(ctx)
	if err != nil {
		return createDTO, err
	}

	spec := rules.RuleAtLeastOneSocialMedia().
		And(rules.RuleValidCityID(citiesIDs)).
		And(rules.RuleValidAdditionalCitiesIDs(citiesIDs)).
		And(rules.RuleValidSpecialityID(specialitiesIDs)).
		And(rules.RuleValidSpecialitiesIDs(specialitiesIDs)).
		And(rules.RuleValidSiteLink()).
		And(rules.RuleValidBirthDate())

	if _, err = spec.IsSatisfied(ctx, &createDTO); err != nil {
		logger.Error(ctx, "Ошибка валидации доктора", err)
		return createDTO, err
	}

	return createDTO, nil
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
