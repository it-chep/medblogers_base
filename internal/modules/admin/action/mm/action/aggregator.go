package action

import (
	"medblogers_base/internal/config"
	"medblogers_base/internal/modules/admin/action/mm/action/change_mm_activity"
	"medblogers_base/internal/modules/admin/action/mm/action/create_getcourse_order"
	"medblogers_base/internal/modules/admin/action/mm/action/create_mm"
	"medblogers_base/internal/modules/admin/action/mm/action/get_mm_list"
	"medblogers_base/internal/modules/admin/action/mm/action/manual_notification_mm"
	"medblogers_base/internal/modules/admin/action/mm/action/push_users_to_mm"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/postgres"
)

type MMActionAggregator struct {
	ChangeMMActivity     *change_mm_activity.Action
	CreateMM             *create_mm.Action
	CreateGetcourceOrder *create_getcourse_order.Action
	GetMMList            *get_mm_list.Action
	ManualNotificationMM *manual_notification_mm.Action
	PushUsersToMM        *push_users_to_mm.Action
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator, config config.AppConfig) *MMActionAggregator {
	return &MMActionAggregator{
		ChangeMMActivity:     change_mm_activity.New(pool),
		CreateMM:             create_mm.New(pool),
		CreateGetcourceOrder: create_getcourse_order.New(pool, clients, config),
		GetMMList:            get_mm_list.New(pool),
		ManualNotificationMM: manual_notification_mm.New(pool, clients),
		PushUsersToMM:        push_users_to_mm.New(pool, clients, config),
	}
}
