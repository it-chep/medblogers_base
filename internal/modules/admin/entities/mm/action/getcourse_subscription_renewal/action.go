package getcourse_subscription_renewal

import (
	"context"
	"fmt"
	"time"

	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/modules/admin/entities/mm/action/getcourse_subscription_renewal/dal"
	"medblogers_base/internal/modules/admin/entities/mm/action/getcourse_subscription_renewal/dto"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
)

type Config interface {
	GetCreateNotificationChatID() int64
}

type Gateway interface {
	NotificateError(ctx context.Context, errText string, clientID int64)
}

type Action struct {
	dal     *dal.Repository
	gateway Gateway
	config  Config
}

func New(pool postgres.PoolWrapper, clients *client.Aggregator, config Config) *Action {
	return &Action{
		dal:     dal.NewRepository(pool),
		gateway: clients.Salebot,
		config:  config,
	}
}

func (a *Action) Do(ctx context.Context, req dto.Request) error {
	user, err := a.dal.GetUserByGKID(ctx, req.GkID)
	if err != nil {
		a.gateway.NotificateError(
			ctx,
			fmt.Sprintf("Ошибка при получении пользователья геткурса %s \n\n data: %v", err.Error(), req),
			a.config.GetCreateNotificationChatID(),
		)
		return err
	}

	return transaction.Exec(ctx, func(ctx context.Context) error {
		if user.ID == 0 {
			userReq := dto.CreateUserRequest{
				GkID:      req.GkID,
				Name:      "",
				EndDate:   time.Now().Add(time.Hour * 24 * time.Duration(req.DaysCount)),
				DaysCount: req.DaysCount,
			}
			err := a.dal.CreateGetcourseUser(ctx, userReq)
			if err != nil {
				a.gateway.NotificateError(
					ctx,
					fmt.Sprintf("Ошибка при сохранении ПОЛЬЗОВАТЕЛЯ от геткурса %s \n\n data: %v", err.Error(), userReq),
					a.config.GetCreateNotificationChatID(),
				)
				return err
			}
		} else {
			allDaysCount := user.DaysCount.Int64 + req.DaysCount

			endTime := time.Now().Add(time.Hour * 24 * time.Duration(req.DaysCount))
			if user.EndDate.Time.After(time.Now()) {
				endTime = user.EndDate.Time.Add(time.Hour * 24 * time.Duration(req.DaysCount))
			}

			err := a.dal.UpdateUserSubscription(ctx, req.GkID, allDaysCount, endTime)
			if err != nil {
				a.gateway.NotificateError(
					ctx,
					fmt.Sprintf(
						"Ошибка при Обновлении подписки пользователя %s \n\n ID: %d, days: %d, time: %s",
						err.Error(),
						req.GkID,
						allDaysCount,
						endTime.Format(time.DateTime),
					),
					a.config.GetCreateNotificationChatID(),
				)
				return err
			}
		}

		return nil
	})
}
