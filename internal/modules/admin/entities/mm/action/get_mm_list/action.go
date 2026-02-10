package get_mm_list

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/mm/action/get_mm_list/dal"
	"medblogers_base/internal/modules/admin/entities/mm/action/get_mm_list/dto"
	"medblogers_base/internal/pkg/postgres"
)

// Action .
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do .
func (a *Action) Do(ctx context.Context) ([]dto.MM, error) {
	return a.dal.GetMMList(ctx)
}
