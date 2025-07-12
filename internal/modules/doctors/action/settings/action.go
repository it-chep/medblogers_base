package settings

import (
	"context"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/action/settings/dal"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/action/settings/dto"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/client"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/dal/city_dal"
	"github.com/it-chep/medblogers_base/internal/modules/doctors/dal/speciality_dal"
	"github.com/it-chep/medblogers_base/internal/pkg/postgres"

	"github.com/it-chep/medblogers_base/internal/modules/doctors/action/settings/service/settings"
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
			dal.NewRepository(pool),
			clients.Subscribers,
		),
	}
}

func (s *Action) Do(ctx context.Context) (dto.Settings, error) {
	err := s.settings.GetSettings(ctx)
	if err != nil {
		return dto.Settings{}, err
	}

	return dto.Settings{}, nil
}
