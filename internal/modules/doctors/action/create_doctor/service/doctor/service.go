package doctor

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/domain/doctor"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/slug"
	"strings"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks . Storage

type Storage interface {
	CreateDoctor(ctx context.Context, createDTO dto.CreateDoctorRequest) (doctor.MedblogersID, error)
	CreateAdditionalCities(ctx context.Context, medblogersID doctor.MedblogersID, citiesIDs []int64) error
	CreateAdditionalSpecialities(ctx context.Context, medblogersID doctor.MedblogersID, specialitiesIDs []int64) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateOrUpdate(ctx context.Context, createDTO dto.CreateDoctorRequest) (dto.CreateDoctorRequest, error) {
	logger.Message(ctx, "[Create] Создание слага и имени")

	createDTO.FullName = s.createName(createDTO.LastName, createDTO.FirstName, createDTO.MiddleName)
	createDTO.Slug = slug.New(createDTO.FullName)
	logger.Message(ctx, "[Create] Сохранение доктора в базе")

	medblogersID, err := s.storage.CreateDoctor(ctx, createDTO)
	if err != nil {
		return dto.CreateDoctorRequest{}, err
	}

	createDTO.ID = medblogersID
	createDTO.AdditionalCities = append(createDTO.AdditionalCities, createDTO.CityID)
	createDTO.AdditionalSpecialties = append(createDTO.AdditionalSpecialties, createDTO.SpecialityID)

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
