package mbc_history

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/mbc_history/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/mbc_history/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Action struct {
	dal *dal.Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64) ([]dto.MBCHistoryItem, error) {
	return a.dal.GetMBCHistory(ctx, doctorID)
}
