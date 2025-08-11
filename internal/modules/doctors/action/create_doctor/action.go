package create_doctor

import (
	"context"
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dal"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/doctor"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/external"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/validate"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/doctors/dal/city_dal"
	"medblogers_base/internal/modules/doctors/dal/speciality_dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action создание врача в базе
type Action struct {
	doctorService     *doctor.Service
	externalService   *external.Service
	validationService *validate.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper, config config.AppConfig) *Action {
	return &Action{
		doctorService:     doctor.NewService(dal.NewRepository(pool)),
		externalService:   external.NewService(clients.Subscribers, clients.Salebot, config),
		validationService: validate.NewService(city_dal.NewRepository(pool), speciality_dal.NewRepository(pool)),
	}
}

func (a *Action) Create(ctx context.Context, createDTO dto.CreateDoctorRequest) ([]dto.ValidationError, error) {
	logger.Message(ctx, fmt.Sprintf("[Create] Создание доктора. Фамилия: %s", createDTO.LastName))

	validationErrors, err := a.validationService.ValidateDoctor(ctx, createDTO)
	if len(validationErrors) > 0 {
		return validationErrors, nil
	}

	_, err = a.doctorService.CreateOrUpdate(ctx, createDTO)
	if err != nil {
		logger.Error(ctx, "Ошибка при сохранении доктора в базе", err)
		return nil, err
	}
	// todo вернуть
	//g := async.NewGroup()
	//g.Go(func() {
	//	a.externalService.NotificatorAdmins(ctx, createdDoctor)
	//})
	//
	//g.Go(func() {
	//	a.externalService.SendToSubscribers(ctx, createdDoctor)
	//})
	//
	//g.Wait()

	return []dto.ValidationError{}, nil
}
