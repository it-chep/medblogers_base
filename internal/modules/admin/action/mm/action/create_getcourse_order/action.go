package create_getcourse_order

import (
	"context"
	"medblogers_base/internal/modules/admin/action/mm/action/create_getcourse_order/dal"
	"medblogers_base/internal/modules/admin/action/mm/action/create_getcourse_order/dto"
	"medblogers_base/internal/pkg/postgres"
	"time"
)

type Action struct {
	dal *dal.Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, req dto.CreateOrderRequest) error {
	orderName, orderDaysCount := a.getNameAndDaysCountFromOrder(req.Positions[0])

	err := a.dal.CreateGetcourseOrder(ctx, dto.GetcourseOrder{
		Name:      orderName,
		DaysCount: orderDaysCount,
		GkID:      req.GkID,
		OrderID:   req.OrderID,
	})
	if err != nil {
		// todo пуш
		return err
	}

	user, err := a.dal.GetUserByGKID(ctx, req.GkID)
	if err != nil {
		err := a.dal.CreateGetcourseUser(ctx, dto.CreateUserRequest{
			GkID:      req.GkID,
			Name:      req.UserName,
			EndDate:   time.Now().Add(time.Hour * 24 * time.Duration(orderDaysCount)),
			DaysCount: orderDaysCount,
		})
		if err != nil {
			// todo пуш
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
			// todo пуш
			return err
		}
	}

	return nil
}

func (a *Action) getNameAndDaysCountFromOrder(orderPosition string) (string, int64) {
	return "", 0 // todo
}
