package action

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/freelancers/action/counters_info"
	"medblogers_base/internal/modules/freelancers/action/create_freelancer"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail"
	"medblogers_base/internal/modules/freelancers/action/get_all_cities"
	"medblogers_base/internal/modules/freelancers/action/get_all_networks"
	"medblogers_base/internal/modules/freelancers/action/get_all_specialities"
	"medblogers_base/internal/modules/freelancers/action/get_pages_count"
	"medblogers_base/internal/modules/freelancers/action/get_seo_detail"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers"
	"medblogers_base/internal/modules/freelancers/action/settings"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	GetSettings            *settings.Action
	GetCounters            *counters_info.Action
	PreliminaryFilterCount *preliminary_filter_count.Action
	GetPagesCount          *get_pages_count.Action

	GetAllCities       *get_all_cities.Action
	GetAllSpecialities *get_all_specialities.Action
	GetAllNetworks     *get_all_networks.Action
	CreateFreelancer   *create_freelancer.Action

	GetSeoDetail     *get_seo_detail.Action
	FreelancerDetail *freelancer_detail.Action

	SearchFreelancers *search_freelancers.Action
	FilterFreelancers *filter_freelancers.Action
}

// NewAggregator конструктор
func NewAggregator(clients *client.Aggregator, pool postgres.PoolWrapper, config config.AppConfig) *Aggregator {
	return &Aggregator{
		GetSettings:            settings.New(pool),
		GetCounters:            counters_info.New(pool),
		PreliminaryFilterCount: preliminary_filter_count.New(pool),
		GetPagesCount:          get_pages_count.New(pool),

		GetAllCities:       get_all_cities.New(pool),
		GetAllSpecialities: get_all_specialities.New(pool),
		GetAllNetworks:     get_all_networks.New(pool),
		CreateFreelancer:   create_freelancer.New(clients, pool, config),

		GetSeoDetail:     get_seo_detail.NewAction(clients, pool),
		FreelancerDetail: freelancer_detail.New(clients, pool),

		SearchFreelancers: search_freelancers.NewAction(clients, pool),
		FilterFreelancers: filter_freelancers.New(clients, pool),
	}
}
