package manual_notification_mm

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/mm/action/manual_notification_mm/dal"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Gateway interface {
	MMNotification(ctx context.Context, clientID int64, mmLink string) error
}

type Action struct {
	dal *dal.Repository
	gw  Gateway
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
		gw:  clients.Salebot,
	}
}

func (a *Action) Do(ctx context.Context, mmID int64) error {
	mm, err := a.dal.GetMMByID(ctx, mmID)
	if err != nil {
		return err
	}

	usersToNotificate, err := a.dal.GetUserToNotificate(ctx)
	if err != nil {
		return err
	}

	newsletterID := uuid.New()
	err = a.dal.CreateNewsletter(ctx, newsletterID, usersToNotificate.GetSbIDs())
	if err != nil {
		return err
	}

	for _, user := range usersToNotificate {
		err = a.gw.MMNotification(ctx, user.SbID.Int64, mm.MMLink.String)
		if err != nil {
			logger.Error(ctx, "send MMNotification error", err)
			continue
		}

		err = a.dal.SendNewsletter(ctx, newsletterID, user.SbID.Int64)
		if err != nil {
			logger.Error(ctx, "save sent newsletter error", err)
			return err
		}
	}

	err = a.dal.SetMMPassed(ctx, mm.ID)
	if err != nil {
		return err
	}

	return nil
}
