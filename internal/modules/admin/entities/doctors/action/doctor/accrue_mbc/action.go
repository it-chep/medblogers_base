package accrue_mbc

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/accrue_mbc/dal"
	pkgctx "medblogers_base/internal/pkg/context"
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

func (a *Action) Do(ctx context.Context, doctorID int64, mbcCount int64) error {
	return a.dal.AccrueMBC(ctx, doctorID, mbcCount, pkgctx.GetUserIDFromContext(ctx))
}
