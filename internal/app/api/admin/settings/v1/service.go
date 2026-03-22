package v1

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/admin"
	"medblogers_base/internal/modules/auth"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/settings/v1"
)

type Implementation struct {
	desc.UnimplementedBannerAdminServiceServer

	admin          *admin.Module
	auth           *auth.Module
	settingsBucket string
}

func New(admin *admin.Module, auth *auth.Module, cfg config.AppConfig) *Implementation {
	return &Implementation{
		admin:          admin,
		auth:           auth,
		settingsBucket: cfg.GetSettingsBucket(),
	}
}
