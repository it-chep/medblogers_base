package preliminary_filter_count

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count/service/freelancers"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_dal"
	domain "medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	freelancersService *freelancers.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		freelancersService: freelancers.NewService(freelancer_dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, filter domain.Filter) (int64, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[PreliminaryFilterCount] Предфильтрация фрилансеров: SocialNetworks: %v, Cities: %v, Specialities: %v, PriceCategory: %v,",
		filter.SocialNetworks, filter.Cities, filter.Specialities, filter.PriceCategory,
	))

	freelancersCount, err := a.freelancersService.GetFreelancersCount(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Ошибка при предфильтрации фрилансеров", err)
		return 0, nil
	}
	return freelancersCount, nil
}
