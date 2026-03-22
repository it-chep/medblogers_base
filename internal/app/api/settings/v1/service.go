package v1

import (
	"medblogers_base/internal/modules/settings"
	desc "medblogers_base/internal/pb/medblogers_base/api/settings/v1"
)

type Implementation struct {
	desc.UnimplementedSettingsServiceServer

	settings *settings.Module
}

func New(module *settings.Module) *Implementation {
	return &Implementation{
		settings: module,
	}
}
