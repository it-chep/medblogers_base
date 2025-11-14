package get_pages_count

import (
	"context"
	"medblogers_base/internal/modules/freelancers/action/get_pages_count/service/page"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_dal"
	"medblogers_base/internal/modules/freelancers/domain/freelancer"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	dal         *freelancer_dal.Repository
	pageService *page.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal:         freelancer_dal.NewRepository(pool),
		pageService: page.New(),
	}
}

func (a *Action) Do(ctx context.Context, filter freelancer.Filter) (int64, error) {
	count, err := a.dal.FreelancersCountByFilter(ctx, filter)
	if err != nil {
		logger.Error(ctx, "ошибка при получении количества фрилансеров по фильру", err)
		return 0, err
	}
	return a.pageService.GetPagesCount(count), nil
}
