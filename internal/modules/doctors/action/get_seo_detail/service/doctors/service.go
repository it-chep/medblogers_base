package doctors

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/doctors/domain/city"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"strings"

	"github.com/samber/lo"
)

// Storage .
type Storage interface {
	GetDoctorInfo(ctx context.Context, slug string) (*doctor.Doctor, error)
	GetDoctorAdditionalCities(ctx context.Context, doctorID doctor.MedblogersID) (map[city.CityID]*city.City, error)
	GetDoctorAdditionalSpecialities(ctx context.Context, doctorID doctor.MedblogersID) (map[speciality.SpecialityID]*speciality.Speciality, error)
}

// Service сервис получения данных о докторе
type Service struct {
	storage Storage
}

// New к-ор
func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// GetDoctorInfo получение информации о докторе
func (s *Service) GetDoctorInfo(ctx context.Context, slug string) (*doctor.Doctor, error) {
	return s.storage.GetDoctorInfo(ctx, slug)
}

func (s *Service) ConfigureDoctorDescription(ctx context.Context, doctorID doctor.MedblogersID) (string, error) {
	var (
		citiesStr       string
		specialitiesStr string
	)
	g := async.NewGroup()

	// получаем города доктора
	g.Go(func() {
		cities, err := s.getCities(ctx, doctorID)
		if err != nil {
			logger.Error(ctx, "[ERROR] Ошибка получения городов для SEO", err)
			return
		}
		citiesStr = cities
	})

	// получаем специальности доктора
	g.Go(func() {
		specialities, err := s.getSpecialities(ctx, doctorID)
		if err != nil {
			logger.Error(ctx, "[ERROR] Ошибка получения специальностей для SEO", err)
			return
		}
		specialitiesStr = specialities
	})

	g.Wait()

	description := fmt.Sprintf("Доктор ведет приемы в городах: %s. Доктор является специалистам по направлениям: %s", citiesStr, specialitiesStr)

	return description, nil
}

func (s *Service) getCities(ctx context.Context, doctorID doctor.MedblogersID) (_ string, err error) {
	citiesMap, err := s.storage.GetDoctorAdditionalCities(ctx, doctorID)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	for i, c := range lo.Values(citiesMap) {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(c.Name())
	}

	return builder.String(), nil
}

func (s *Service) getSpecialities(ctx context.Context, doctorID doctor.MedblogersID) (_ string, err error) {
	specialitiesMap, err := s.storage.GetDoctorAdditionalSpecialities(ctx, doctorID)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	for i, sp := range lo.Values(specialitiesMap) {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(sp.Name())
	}

	return builder.String(), nil
}
