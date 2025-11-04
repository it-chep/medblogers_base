package freelancer

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/modules/freelancers/domain/speciality"

	"github.com/samber/lo"

	"medblogers_base/internal/pkg/logger"
	"strings"
)

// Storage .
type Storage interface {
	GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error)
	GetFreelancerAdditionalCities(ctx context.Context, freelancerID int64) (map[int64]*city.City, error)
	GetFreelancerAdditionalSpecialities(ctx context.Context, freelancerID int64) (map[int64]*speciality.Speciality, error)
}

type ImageGetter interface {
	GetPhotoLink(s3Key string) string
}

// Service сервис получения данных о докторе
type Service struct {
	storage     Storage
	imageGetter ImageGetter
}

// New к-ор
func New(storage Storage, imageGetter ImageGetter) *Service {
	return &Service{
		storage:     storage,
		imageGetter: imageGetter,
	}
}

// GetFreelancerInfo получение информации о докторе
func (s *Service) GetFreelancerInfo(ctx context.Context, slug string) (*freelancer.Freelancer, error) {
	return s.storage.GetFreelancerInfo(ctx, slug)
}

func (s *Service) ConfigureFreelancerDescription(ctx context.Context, frlncr *freelancer.Freelancer) (string, error) {
	var (
		specialitiesStr string
	)

	specialitiesStr, err := s.getSpecialities(ctx, frlncr.GetID())
	if err != nil {
		logger.Error(ctx, "[ERROR] Ошибка получения специальностей для SEO", err)
		return "", err
	}

	description := fmt.Sprintf(
		"Роли: %s. Опыт работы %s. %s",
		specialitiesStr,
		frlncr.GetWorkingExperience(),
		lo.Ternary(frlncr.HasExperienceWithDoctor(), "Есть опыт работы с врачами.", ""),
	)

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

// GetFreelancerPhoto получение фотки фрилансера
func (s *Service) GetFreelancerPhoto(s3Key string) string {
	return s.imageGetter.GetPhotoLink(s3Key)
}
