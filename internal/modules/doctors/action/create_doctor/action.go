package create_doctor

import (
	"context"
	"fmt"
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dal"
	"medblogers_base/internal/modules/doctors/action/create_doctor/dto"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/doctor"
	"medblogers_base/internal/modules/doctors/action/create_doctor/service/external"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/pipe"
	"medblogers_base/internal/pkg/postgres"
)

// Action создание врача в базе
type Action struct {
	doctorService   *doctor.Service
	externalService *external.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper, config *config.Config) *Action {
	return &Action{
		doctorService:   doctor.NewService(dal.NewRepository(pool)),
		externalService: external.NewService(clients.Subscribers, clients.Salebot, clients, config),
	}
}

func (a *Action) Create(ctx context.Context, createDTO dto.CreateDoctorRequest) error {
	logger.Message(ctx, fmt.Sprintf("[Create] Создание доктора. Фамилия: %s", createDTO.LastName))

	errors, err := pipe.With(a.doctorService.ValidateDoctor).
		With(a.doctorService.CreateOrUpdate).
		With(a.externalService.NotificatorAdmins).
		With(a.externalService.SendToSubscribers).
		Run(ctx, createDTO).Get()
}
