package activate

import (
	"context"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/activate/dal"
	common_dal "medblogers_base/internal/modules/admin/entities/doctors/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/doctor"
	"medblogers_base/internal/pkg/postgres"
)

type Subscribers interface {
	ActivateDoctor(ctx context.Context, doctorID int64) error
}

type CommonDal interface {
	GetDoctorByID(ctx context.Context, doctorID int64) (*doctor.Doctor, error)
}

type ActionDal interface {
	ActivateDoctor(ctx context.Context, doctorID int64) (err error)
}

// Action активация доктора
type Action struct {
	commonDal   CommonDal
	actionDal   ActionDal
	subscribers Subscribers
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		actionDal:   dal.NewRepository(pool),
		commonDal:   common_dal.NewRepository(pool),
		subscribers: clients.Subscribers,
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64) error {
	doc, err := a.commonDal.GetDoctorByID(ctx, doctorID)
	if err != nil {
		return err
	}

	if doc.GetIsActive() {
		return errors.New("Доктор уже активен")
	}

	err = a.subscribers.ActivateDoctor(ctx, doctorID)
	if err != nil {
		return err
	}

	return a.actionDal.ActivateDoctor(ctx, doctorID)
}
