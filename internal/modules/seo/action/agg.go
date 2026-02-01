package action

import (
	"medblogers_base/internal/modules/seo/action/get_sitemap_info"
	"medblogers_base/internal/pkg/postgres"
)

type Aggregator struct {
	GetSitemapInfo *get_sitemap_info.Action
}

func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		GetSitemapInfo: get_sitemap_info.New(pool),
	}
}
