package actions

import (
	"medblogers_base/internal/modules/analytics/actions/create_cookie_id"
	"medblogers_base/internal/modules/analytics/actions/save_analytics"
	"medblogers_base/internal/modules/analytics/actions/save_site_form_answer"
	"medblogers_base/internal/modules/analytics/actions/update_cookie_activity"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает action-ы модуля аналитики.
type Aggregator struct {
	CreateCookieID       *create_cookie_id.Action
	SaveSiteFormAnswer   *save_site_form_answer.Action
	SaveAnalytics        *save_analytics.Action
	UpdateCookieActivity *update_cookie_activity.Action
}

// NewAggregator .
func NewAggregator(pool postgres.PoolWrapper) *Aggregator {
	return &Aggregator{
		CreateCookieID:       create_cookie_id.New(pool),
		SaveSiteFormAnswer:   save_site_form_answer.New(pool),
		SaveAnalytics:        save_analytics.New(pool),
		UpdateCookieActivity: update_cookie_activity.New(pool),
	}
}
