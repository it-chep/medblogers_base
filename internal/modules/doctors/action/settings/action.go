package settings

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/settings/service/settings"
)

// Action получение настроек главной страницы
type Action struct {
	settings *settings.Service
}

// New .
func New() *Action {
	return &Action{
		settings: settings.NewSettingsService(),
	}
}

func (s *Action) GetSettings(ctx context.Context) error {
	return s.settings.GetSettings(ctx)
}
