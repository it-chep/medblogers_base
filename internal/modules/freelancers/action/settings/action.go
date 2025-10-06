package settings

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/settings/dal"
	"medblogers_base/internal/modules/freelancers/action/settings/dto"
	"medblogers_base/internal/modules/freelancers/action/settings/service/settings"
	"medblogers_base/internal/modules/freelancers/dal/city_dal"
	"medblogers_base/internal/modules/freelancers/dal/society_dal"
	"medblogers_base/internal/modules/freelancers/dal/speciality_dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение настроек главной страницы
type Action struct {
	settings *settings.Service
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		settings: settings.NewSettingsService(
			city_dal.NewRepository(pool),
			speciality_dal.NewRepository(pool),
			society_dal.NewRepository(pool),
			dal.NewRepository(pool),
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
