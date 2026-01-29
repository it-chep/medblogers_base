package push_users_to_mm

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/action/mm/action/push_users_to_mm/dal"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Config interface {
	GetCreateNotificationChatID() int64
}

type Gateway interface {
	MMNotification(ctx context.Context, clientID int64, mmLink string) error
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
	mm, err := a.dal.GetNearestMM(ctx)
	if err != nil {
		logger.Error(ctx, "get MM error", err)
		return err
	}

	if mm.ID == 0 {
		return nil
	}

	usersToNotificate, err := a.dal.GetUserToNotificate(ctx)
	if err != nil {
		logger.Error(ctx, "get users error", err)
		return err
	}

	newsletterID := uuid.New()
	err = a.dal.CreateNewsletter(ctx, newsletterID, usersToNotificate.GetSbIDs())
	if err != nil {
		logger.Error(ctx, "create newsletter error", err)
		return err
	}

	for _, user := range usersToNotificate {
		err = a.gw.MMNotification(ctx, user.SbID.Int64, mm.MMLink.String)
		if err != nil {
			logger.Error(ctx, "send MMNotification error", err)
			a.gw.NotificateError(ctx, fmt.Sprintf("Ошибка при отправки сообщеня %s \n\n err: %s \n\n SB_ID: %d", newsletterID.String(), err.Error(), user.SbID.Int64), a.config.GetCreateNotificationChatID())
			continue
		}

		err = a.dal.SendNewsletter(ctx, newsletterID, user.SbID.Int64)
		if err != nil {
			logger.Error(ctx, "save sent newsletter error", err)
			a.gw.NotificateError(ctx, fmt.Sprintf("Ошибка при сохранении рассылки %s \n\n err: %s", newsletterID.String(), err.Error()), a.config.GetCreateNotificationChatID())
		}
	}

	err = a.dal.SetMMPassed(ctx, mm.ID)
	if err != nil {
		logger.Error(ctx, "SetMMPassed error", err)
		return err
	}

	return nil
}
