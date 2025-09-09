package freelancer

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/speciality"

	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"strings"
)

// Storage .
type Storage interface {
	GetFreelancerInfo(ctx context.Context, slug string) (*doctor.Doctor, error)
	GetFreelancerAdditionalCities(ctx context.Context, freelancerID int64) (map[int64]*city.City, error)
	GetFreelancerAdditionalSpecialities(ctx context.Context, freelancerID int64) (map[int64]*speciality.Speciality, error)
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

// GetFreelancerInfo получение информации о докторе
func (s *Service) GetFreelancerInfo(ctx context.Context, slug string) (*doctor.Doctor, error) {
	return s.storage.GetFreelancerInfo(ctx, slug)
}

func (s *Service) ConfigureDoctorDescription(ctx context.Context, freelancerID int64) (string, error) {
	var (
		citiesStr       string
		specialitiesStr string
	)
	g := async.NewGroup()

	// получаем города фрилансера
	g.Go(func() {
		cities, err := s.getCities(ctx, freelancerID)
		if err != nil {
			logger.Error(ctx, "[ERROR] Ошибка получения городов для SEO", err)
			return
		}
		citiesStr = cities
	})

	// получаем специальности фрилансера
	g.Go(func() {
		specialities, err := s.getSpecialities(ctx, freelancerID)
		if err != nil {
			logger.Error(ctx, "[ERROR] Ошибка получения специальностей для SEO", err)
			return
		}
		specialitiesStr = specialities
	})

	g.Wait()

	description := fmt.Sprintf("Работаю в городах: %s. Яляюсь специалистам по направлениям: %s", citiesStr, specialitiesStr)

	return description, nil
}

func (s *Service) getCities(ctx context.Context, freelancerID int64) (_ string, err error) {
	citiesMap, err := s.storage.GetFreelancerAdditionalCities(ctx, freelancerID)
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

func (s *Service) getSpecialities(ctx context.Context, freelancerID int64) (_ string, err error) {
	specialitiesMap, err := s.storage.GetFreelancerAdditionalSpecialities(ctx, freelancerID)
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
