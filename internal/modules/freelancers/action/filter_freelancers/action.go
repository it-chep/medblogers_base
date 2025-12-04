package filter_freelancers

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers/dal"
	"medblogers_base/internal/modules/freelancers/action/filter_freelancers/dto"
	filterService "medblogers_base/internal/modules/freelancers/action/filter_freelancers/service/freelancer"
	"medblogers_base/internal/modules/freelancers/client"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action фильтрация фрилансеров
type Action struct {
	filterService *filterService.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	repository := dal.NewRepository(pool)
	return &Action{
		filterService: filterService.NewService(repository, clients.S3, repository),
	}
}

// Do фильтрация фрилансеров
func (a *Action) Do(ctx context.Context, filter freelancer.Filter) ([]dto.Freelancer, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[Filter] Фильтрация фрилансеров, Cities: %v, Specialities: %v, Page: %d, ",
		filter.Cities, filter.Specialities, filter.Page,
	))

	if len(filter.Cities) != 0 ||
		len(filter.Specialities) != 0 ||
		len(filter.SocialNetworks) != 0 ||
		len(filter.PriceCategory) != 0 {
		return a.filterService.GetFreelancersByFilter(ctx, filter)
	}

	return a.filterService.GetFreelancers(ctx, filter)
}
