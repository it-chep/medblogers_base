package preliminary_filter_count

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count/dal"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count/dto"
	"medblogers_base/internal/modules/freelancers/action/preliminary_filter_count/service/freelancers"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	freelancersService *freelancers.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		freelancersService: freelancers.NewService(dal.NewRepository(pool)),
	}
}

func (a *Action) Do(ctx context.Context, filter dto.Filter) (int64, error) {
	logger.Message(ctx, fmt.Sprintf(
		"[PreliminaryFilterCount] Предфильтрация фрилансеров: SocialNetworks: %v, Cities: %v, Specialities: %v, WorkDoc: %t, PriceCategory: %v,",
		filter.SocialNetworks, filter.Cities, filter.Specialities, filter.ExperienceWithDoctors, filter.PriceCategory,
	))

	freelancersCount, err := a.freelancersService.GetFreelancersCount(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Ошибка при предфильтрации фрилансеров", err)
		return 0, nil
	}
	return freelancersCount, nil
}
