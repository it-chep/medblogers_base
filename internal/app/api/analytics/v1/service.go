package v1

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/analytics"
	desc "medblogers_base/internal/pb/medblogers_base/api/analytics/v1"
)

type Implementation struct {
	desc.UnimplementedAnalyticsServiceServer

	analytics *analytics.Module
	config    config.AppConfig
}

// NewService .
func NewService(module *analytics.Module, cfg config.AppConfig) *Implementation {
	return &Implementation{
		analytics: module,
		config:    cfg,
	}
}
