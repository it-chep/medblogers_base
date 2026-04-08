package actions

import (
	"medblogers_base/internal/modules/analytics/actions/create_cookie_id"
	"medblogers_base/internal/modules/analytics/actions/update_cookie_activity"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает action-ы модуля аналитики.
type Aggregator struct {
	CreateCookieID       *create_cookie_id.Action
	UpdateCookieActivity *update_cookie_activity.Action
}

// NewAggregator .
func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		CreateCookieID:       create_cookie_id.New(pool),
		UpdateCookieActivity: update_cookie_activity.New(pool),
	}
}
