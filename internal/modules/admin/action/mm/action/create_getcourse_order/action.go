package create_getcourse_order

import (
	"context"
	"fmt"
	"medblogers_base/internal/modules/admin/action/mm/action/create_getcourse_order/dal"
	"medblogers_base/internal/modules/admin/action/mm/action/create_getcourse_order/dto"
	"medblogers_base/internal/modules/admin/client"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/transaction"
	"strings"
	"time"
)

type Config interface {
	GetCreateNotificationChatID() int64
}

// Gateway ..
type Gateway interface {
	NotificateError(ctx context.Context, errText string, clientID int64)
}

// Action ..
type Action struct {
	dal     *dal.Repository
	gateway Gateway
	config  Config
}

// New ..
func New(pool postgres.PoolWrapper, clients *client.Aggregator, config Config) *Action {
	return &Action{
		dal:     dal.NewRepository(pool),
		gateway: clients.Salebot,
		config:  config,
	}
}

// Do .
func (a *Action) Do(ctx context.Context, req dto.CreateOrderRequest) error {
	orderName, orderDaysCount := a.getNameAndDaysCountFromOrder(req.Position)
	if orderDaysCount == 0 {
		return nil
	}

	user, err := a.dal.GetUserByGKID(ctx, req.GkID)
	if err != nil {
		a.gateway.NotificateError(ctx, fmt.Sprintf("Ошибка при получении пользователья геткурса %s \n\n data: %v", err.Error(), req), a.config.GetCreateNotificationChatID())
		return err
	}

	logger.Message(ctx, fmt.Sprintf("Пришел заказ от ГК: %v", req))

	return transaction.Exec(ctx, func(ctx context.Context) error {
		err = a.dal.CreateGetcourseOrder(ctx, dto.GetcourseOrder{
			Name:      orderName,
			DaysCount: orderDaysCount,
			GkID:      req.GkID,
			OrderID:   req.OrderID,
		})
		if err != nil {
			a.gateway.NotificateError(ctx, fmt.Sprintf("Ошибка при сохранении ЗАКАЗА от геткурса %s \n\n data: %v", err.Error(), req), a.config.GetCreateNotificationChatID())
			return err
		}

		if user.ID == 0 {
			userReq := dto.CreateUserRequest{
				GkID:      req.GkID,
				Name:      req.UserName,
				EndDate:   time.Now().Add(time.Hour * 24 * time.Duration(orderDaysCount)),
				DaysCount: orderDaysCount,
			}
			err := a.dal.CreateGetcourseUser(ctx, userReq)
			if err != nil {
				a.gateway.NotificateError(ctx, fmt.Sprintf("Ошибка при сохранении ПОЛЬЗОВАТЕЛЯ от геткурса %s \n\n data: %v", err.Error(), userReq), a.config.GetCreateNotificationChatID())
				return err
			}
		} else {
			allDaysCount := user.DaysCount.Int64 + orderDaysCount

			endTime := time.Now().Add(time.Hour * 24 * time.Duration(orderDaysCount))

			if user.EndDate.Time.After(time.Now()) {
				// Если подписка пока не заканчивается, то добавляем дни к общему концу подписки
				endTime = user.EndDate.Time.Add(time.Hour * 24 * time.Duration(orderDaysCount))
			}
			err := a.dal.UpdateUserSubscription(ctx, req.GkID, allDaysCount, endTime)
			if err != nil {
				a.gateway.NotificateError(
					ctx,
					fmt.Sprintf("Ошибка при Обновлении подписки пользователя %s \n\n ID: %d, days: %d, time: %s",
						err.Error(), req.GkID, allDaysCount, endTime.Format(time.DateTime),
					),
					a.config.GetCreateNotificationChatID(),
				)
				return err
			}
		}
		return nil
	})
}

func (a *Action) getNameAndDaysCountFromOrder(orderPosition string) (string, int64) {
	infoMap := map[string]int64{
		"1 месяц":    30,
		"2 месяца":   60,
		"3 месяца":   90,
		"6 месяцев":  180,
		"12 месяцев": 365,
	}

	daysCount := int64(0)
	for k, v := range infoMap {
		if strings.Contains(orderPosition, k) {
			daysCount = v
			break
		}
	}

	return orderPosition, daysCount
}
