package check_sb_id

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/admin/action/mm/action/check_sb_id/dal"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/postgres"
)

type Config interface {
	GetCreateNotificationChatID() int64
}

type Gateway interface {
	NotificateError(ctx context.Context, errText string, clientID int64)
}

type Action struct {
	dal    *dal.Repository
	gw     Gateway
	config Config
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator, config Config) *Action {
	return &Action{
		dal:    dal.NewRepository(pool),
		gw:     clients.Salebot,
		config: config,
	}
}

func (a *Action) Do(ctx context.Context) error {
	adminID := a.config.GetCreateNotificationChatID()

	users, err := a.dal.GetEmptySBIDUsers(ctx)
	if len(users) == 0 {
		return nil
	}

	if err != nil {
		a.gw.NotificateError(ctx, fmt.Sprintf("Ошибка при получении пользаков без SB_ID %s", err), adminID)
		return nil
	}

	notificationText := fmt.Sprintf("Найдены пользователи из геткурса без sb_id: \n\n GK IDs - %s", users.StringGkIDs())

	a.gw.NotificateError(ctx, notificationText, adminID)
	return nil
}
