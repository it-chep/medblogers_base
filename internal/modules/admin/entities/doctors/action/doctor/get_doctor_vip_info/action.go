package get_doctor_vip_info

import (
	"context"
	common_dal "medblogers_base/internal/modules/admin/entities/doctors/dal/vip_card_dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/vip_card"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	GetVipCardInfo(ctx context.Context, doctorID int64) (*vip_card.VipCard, error)
}

type Action struct {
	dal Dal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: common_dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64) (*vip_card.VipCard, error) {
	return a.dal.GetVipCardInfo(ctx, doctorID)
}
