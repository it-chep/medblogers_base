package settings

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/settings/dto"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/doctors/dal/city_dal"
	"medblogers_base/internal/modules/doctors/dal/speciality_dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"

	"medblogers_base/internal/modules/doctors/action/settings/service/settings"
)

// Action получение настроек главной страницы
type Action struct {
	settings *settings.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		settings: settings.NewSettingsService(
			city_dal.NewRepository(pool),
			speciality_dal.NewRepository(pool),
			clients.Subscribers,
		),
	}
}

func (s *Action) Do(ctx context.Context) (*dto.Settings, error) {
	logger.Message(ctx, "[Settings] Получение настроек для главной страницы")

	settingsDTO, err := s.settings.GetSettings(ctx)
	if err != nil {
		return &dto.Settings{}, err
	}

	return settingsDTO, nil
}
