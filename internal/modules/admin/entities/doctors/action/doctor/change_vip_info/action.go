package change_vip_info

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_vip_info/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_vip_info/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_vip_info/service/audit"
	common_dal "medblogers_base/internal/modules/admin/entities/doctors/dal/vip_card_dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/vip_card"
	"medblogers_base/internal/pkg/audit_logger"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type Dal interface {
	UPSERTVipInfo(ctx context.Context, doctorID int64, updateVipRequest dto.UpdateRequest) error
}

type CommonDal interface {
	GetVipCardInfo(ctx context.Context, doctorID int64) (*vip_card.VipCard, error)
}

// Action обновление випки
type Action struct {
	dal       Dal
	commonDal CommonDal
	audit     *audit.Service
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal:       dal.NewRepository(pool),
		commonDal: common_dal.NewRepository(pool),
		audit:     audit.New(audit_logger.NewLogger(pool)),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64, updateReq dto.UpdateRequest) error {
	vipCard, err := a.commonDal.GetVipCardInfo(ctx, doctorID)
	if err != nil {
		return err
	}
	return transaction.Exec(ctx, func(ctx context.Context) error {
		err = a.audit.Log(ctx, vipCard)
		if err != nil {
			return err
		}

		return a.dal.UPSERTVipInfo(ctx, doctorID, updateReq)
	})
}
