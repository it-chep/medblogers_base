package check_expired_vip_cards

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_doctor_vip_activity"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/check_expired_vip_cards/dal"
	"medblogers_base/internal/pkg/postgres"
)

type Config interface {
	GetCreateNotificationChatID() int64
}

type Dal interface {
	GetExpiredVipDoctorIDs(ctx context.Context) ([]int64, error)
}

type Gateway interface {
	NotificateError(ctx context.Context, errText string, clientID int64)
}

type VipActivity interface {
	Do(ctx context.Context, doctorID int64, activity bool) error
}

type Action struct {
	config      Config
	gateway     Gateway
	dal         Dal
	vipActivity VipActivity
}

func New(clients *client.Aggregator, pool postgres.PoolWrapper, config Config) *Action {
	return &Action{
		config:      config,
		gateway:     clients.Salebot,
		dal:         dal.NewRepository(pool),
		vipActivity: change_doctor_vip_activity.New(clients, pool),
	}
}

func (a *Action) Do(ctx context.Context) error {
	doctorIDs, err := a.dal.GetExpiredVipDoctorIDs(ctx)
	if err != nil {
		a.gateway.NotificateError(
			ctx,
			fmt.Sprintf("Ошибка в джобе check_expired_vip_cards при получении докторов: %s", err.Error()),
			a.config.GetCreateNotificationChatID(),
		)
		return err
	}

	for _, doctorID := range doctorIDs {
		if err = a.vipActivity.Do(ctx, doctorID, false); err != nil {
			a.gateway.NotificateError(
				ctx,
				fmt.Sprintf("Ошибка в джобе check_expired_vip_cards при деактивации VIP у доктора %d: %s", doctorID, err.Error()),
				a.config.GetCreateNotificationChatID(),
			)
			return err
		}
	}

	return nil
}
