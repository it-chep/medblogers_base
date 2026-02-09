package create_mm

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/mm/action/create_mm/dal"
	"medblogers_base/internal/modules/admin/entities/mm/action/create_mm/dto"
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
func (a *Action) Do(ctx context.Context, req dto.CreateMMRequest) error {
	return a.dal.CreateMM(ctx, req)
}
