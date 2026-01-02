package get

import (
	"context"
	common_dal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
)

type CommonDal interface {
	GetDoctors(ctx context.Context) ([]*doctor.Doctor, error)
}

// Action список докторов
type Action struct {
	commonDal CommonDal
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: common_dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context) ([]*doctor.Doctor, error) {
	return a.commonDal.GetDoctors(ctx)
}
