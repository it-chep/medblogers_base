package v1

import (
	"medblogers_base/internal/modules/analytics"
	desc "medblogers_base/internal/pb/medblogers_base/api/analytics/v1"
)

type Implementation struct {
	desc.UnimplementedAnalyticsServiceServer

	analytics *analytics.Module
}

// NewService .
func NewService(module *analytics.Module) *Implementation {
	return &Implementation{
		analytics: module,
	}
}
