package action

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers"
	"medblogers_base/internal/modules/freelancers/action/freelancer_detail"
	"medblogers_base/internal/modules/freelancers/action/get_all_cities"
	"medblogers_base/internal/modules/freelancers/action/get_all_networks"
	"medblogers_base/internal/modules/freelancers/action/get_all_specialities"
	"medblogers_base/internal/modules/freelancers/action/get_pages_count"
	"medblogers_base/internal/modules/freelancers/action/get_seo_detail"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count"
	"medblogers_base/internal/modules/freelancers/action/search_freelancers"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/pkg/postgres"
)

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	GetPagesCount      *get_pages_count.Action
	GetAllCities       *get_all_cities.Action
	GetAllSpecialities *get_all_specialities.Action
	GetAllNetworks     *get_all_networks.Action
	GetSeoDetail       *get_seo_detail.Action

	FreelancerDetail       *freelancer_detail.Action
	SearchFreelancers      *search_freelancers.Action
	FilterFreelancers      *filter_freelancers.Action
	PreliminaryFilterCount *preliminary_filter_count.Action
}

// NewAggregator конструктор
func NewAggregator(clients *client.Aggregator, pool postgres.PoolWrapper, config config.AppConfig) *Aggregator {
	return &Aggregator{
		PreliminaryFilterCount: preliminary_filter_count.New(pool),
	}
}
