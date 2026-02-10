package update

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/dal"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/dto"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/service/audit"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/service/freelancer"
	commondal "medblogers_base/internal/modules/admin/entities/freelancers/dal"
	"medblogers_base/internal/pkg/audit_logger"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type Action struct {
	audit      *audit.Service
	freelancer *freelancer.Service
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		audit: audit.New(audit_logger.NewLogger(pool)),
		freelancer: freelancer.NewService(
			commondal.NewRepository(pool),
			dal.NewRepository(pool),
		),
	}
}

// Do обновление врача
func (a *Action) Do(ctx context.Context, doctorID int64, updateReq dto.UpdateRequest) error {
	frlncr, err := a.freelancer.GetFreelancer(ctx, doctorID)
	if err != nil {
		return err
	}
	return transaction.Exec(ctx, func(ctx context.Context) error {
		err = a.audit.Log(ctx, frlncr)
		if err != nil {
			return err
		}

		return a.freelancer.UpdateFreelancer(ctx, doctorID, updateReq)
	})
}
