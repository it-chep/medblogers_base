package action

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/settings/action/get_banners"
	"medblogers_base/internal/pkg/postgres"
)

type Aggregator struct {
	GetBanners *get_banners.Action
}

func New(pool postgres.PoolWrapper, cfg config.AppConfig) *Aggregator {
	return &Aggregator{
		GetBanners: get_banners.New(pool, cfg),
	}
}
