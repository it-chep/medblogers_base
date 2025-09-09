package counters_info

import (
	"context"
	"medblogers_base/internal/modules/freelancers/dal/freelancer_dal"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение настроек главной страницы
type Action struct {
	dal *freelancer_dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: freelancer_dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) (int64, error) {
	return a.dal.GetFreelancersCount(ctx)
}
