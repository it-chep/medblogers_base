package counters_info

import (
	"context"
	"medblogers_base/internal/modules/doctors/action/counters_info/dto"
	"medblogers_base/internal/modules/doctors/action/counters_info/service/counters"
	"medblogers_base/internal/modules/doctors/client"
	"medblogers_base/internal/modules/doctors/dal/doctor_dal"
	"medblogers_base/internal/pkg/async"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

// Action получение настроек главной страницы
type Action struct {
	counters *counters.Service
}

// New .
func New(clients *client.Aggregator, pool postgres.PoolWrapper) *Action {
	return &Action{
		counters: counters.NewService(
			doctor_dal.NewRepository(pool),
			clients.Subscribers,
		),
	}
}

func (s *Action) Do(ctx context.Context) (dto.CountersInfoDTO, error) {
	logger.Message(ctx, "[Counters] Получение общего количества докторов и подписчиков")
	var (
		err              error
		doctorsCount     int64
		subscribersCount string
	)
	g := async.NewGroup()

	g.Go(func() {
		doctorsCount, err = s.counters.GetDoctorsCount(ctx)
		if err != nil {
			logger.Error(ctx, "[Counters] Ошибка получения докторов", err)
		}
	})

	g.Go(func() {
		subscribersCount, err = s.counters.GetSubscribersCount(ctx)
		if err != nil {
			logger.Error(ctx, "[Counters] Ошибка получения подписчиков", err)
		}
	})

	g.Wait()

	return dto.CountersInfoDTO{
		DoctorsCount:     doctorsCount,
		SubscribersCount: subscribersCount,
	}, nil
}
