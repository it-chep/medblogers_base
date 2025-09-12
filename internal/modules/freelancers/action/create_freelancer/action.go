package create_freelancer

import (
	"context"
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dal"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/dto"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/service/external"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/service/freelancer"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer/service/validate"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/modules/freelancers/dal/city_dal"
	"medblogers_base/internal/modules/freelancers/dal/speciality_dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	creationService   *freelancer.Service
	externalService   *external.Service
	validationService *validate.Service
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper, config config.AppConfig) *Action {
	return &Action{
		creationService:   freelancer.NewService(dal.NewRepository(pool)),
		externalService:   external.NewService(clients.Salebot, config),
		validationService: validate.NewService(city_dal.NewRepository(pool), speciality_dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, createDTO dto.CreateRequest) ([]dto.ValidationError, error) {
	logger.Message(ctx, fmt.Sprintf("[Create] Создание ФРИЛАНСЕРА. Фамилия: %s", createDTO.LastName))
	logger.Message(ctx, fmt.Sprintf(
		"[Create] Поля: CityID=%d, SpecialityID=%d, Email=%s, LastName=%s, FirstName=%s, MiddleName=%s, AdditionalCities=%v, AdditionalSpecialties=%v",
		createDTO.MainCityID, createDTO.MainSpecialityID, createDTO.Email, createDTO.LastName, createDTO.FirstName,
		createDTO.MiddleName,
		createDTO.AdditionalCities, createDTO.AdditionalSpecialties,
	))

	validationErrors, err := a.validationService.ValidateFreelancer(ctx, &createDTO)
	if len(validationErrors) > 0 {
		return validationErrors, nil
	}

	_, err = a.creationService.CreateOrUpdate(ctx, createDTO)
	if err != nil {
		logger.Error(ctx, "Ошибка при сохранении фрилансера в базе", err)
		return nil, err
	}

	a.externalService.NotificatorAdmins(ctx, createDTO)

	return []dto.ValidationError{}, nil
}
