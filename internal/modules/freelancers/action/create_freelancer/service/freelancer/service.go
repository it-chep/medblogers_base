package freelancer

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/slug"
	"strings"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	CreateFreelancer(ctx context.Context, createDTO dto.CreateRequest) (int64, error)
	CreateAdditionalCities(ctx context.Context, medblogersID int64, citiesIDs []int64) error
	CreateAdditionalSpecialities(ctx context.Context, medblogersID int64, specialitiesIDs []int64) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateOrUpdate(ctx context.Context, createDTO dto.CreateRequest) (dto.CreateRequest, error) {
	logger.Message(ctx, "[Create] Создание слага и имени")

	createDTO.Name = s.createName(createDTO.LastName, createDTO.FirstName, createDTO.MiddleName)
	createDTO.Slug = slug.New(createDTO.Name)
	logger.Message(ctx, "[Create] Сохранение фрилансера в базе")

	medblogersID, err := s.storage.CreateFreelancer(ctx, createDTO)
	if err != nil {
		return dto.CreateRequest{}, err
	}

	createDTO.ID = medblogersID
	createDTO.AdditionalCities = append(createDTO.AdditionalCities, createDTO.MainCityID)
	createDTO.AdditionalSpecialties = append(createDTO.AdditionalSpecialties, createDTO.MainSpecialityID)

	g := async.NewGroup()

	logger.Message(ctx, "[Create] Сохранение дополнительных параметров")
	g.Go(func() {
		err = s.storage.CreateAdditionalCities(ctx, medblogersID, createDTO.AdditionalCities)
		if err != nil {
			logger.Error(ctx, "[Create] Ошибка при сохранении доп городов ", err)
		}
	})
	g.Go(func() {
		err = s.storage.CreateAdditionalSpecialities(ctx, medblogersID, createDTO.AdditionalSpecialties)
		if err != nil {
			logger.Error(ctx, "[Create] Ошибка при сохранении доп специальностей ", err)
		}
	})

	g.Wait()

	return createDTO, nil
}

func (s *Service) createName(lastName, firstName, middleName string) string {
	parts := []string{
		strings.TrimSpace(lastName),
		strings.TrimSpace(firstName),
		strings.TrimSpace(middleName),
	}

	// Удаляем пустые части
	var nonEmptyParts []string
	for _, part := range parts {
		if part != "" {
			nonEmptyParts = append(nonEmptyParts, part)
		}
	}

	return strings.Join(nonEmptyParts, " ")
}
