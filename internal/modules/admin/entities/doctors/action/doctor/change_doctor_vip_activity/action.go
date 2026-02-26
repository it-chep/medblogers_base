package change_doctor_vip_activity

import (
	"context"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_doctor_vip_activity/dal"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_doctor_vip_activity/service/subscribers"
	"medblogers_base/internal/pkg/postgres"
)

type Dal interface {
	ChangeDoctorVipActivity(ctx context.Context, doctorID int64, activity bool) error
}

type Action struct {
	subscribers *subscribers.Service
	dal         Dal
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		subscribers: subscribers.New(clients.Subscribers),
		dal:         dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, doctorID int64, activity bool) error {
	err := a.subscribers.ChangeVipActivity(ctx, doctorID, activity)
	if err != nil {
		return err
	}

	return a.dal.ChangeDoctorVipActivity(ctx, doctorID, activity)
}
