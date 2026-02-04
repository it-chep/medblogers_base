package update

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/dto"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/service/audit"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/update/service/doctor"
	commondal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/pkg/audit_logger"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type Action struct {
	audit  *audit.Service
	doctor *doctor.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		audit: audit.New(audit_logger.NewLogger(pool)),
		doctor: doctor.NewService(
			commondal.NewRepository(pool),
			dal.NewRepository(pool),
		),
	}
}

// Do обновление врача
func (a *Action) Do(ctx context.Context, doctorID int64, updateReq dto.UpdateRequest) error {
	doc, err := a.doctor.GetDoctor(ctx, doctorID)
	if err != nil {
		return err
	}
	return transaction.Exec(ctx, func(ctx context.Context) error {
		err = a.audit.Log(ctx, doc)
		if err != nil {
			return err
		}

		return a.doctor.UpdateDoctor(ctx, doctorID, updateReq)
	})
}
