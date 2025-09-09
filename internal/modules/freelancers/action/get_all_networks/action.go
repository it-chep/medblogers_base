package get_all_networks

import (
	"context"
	"medblogers_base/internal/modules/freelancers/dal/society_dal"
	"medblogers_base/internal/modules/freelancers/domain/social_network"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Storage interface {
	GetAllNetworks(ctx context.Context) ([]*social_network.SocialNetwork, error)
}

// Action список соц сетей в для регистрации
type Action struct {
	storage Storage
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		storage: society_dal.NewRepository(pool),
	}
}

// Do выполнение
func (a Action) Do(ctx context.Context) ([]*social_network.SocialNetwork, error) {
	logger.Message(ctx, "[Reg] Получение соц сетей для регистрации")
	networks, err := a.storage.GetAllNetworks(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка при получении соц сетей для регистрации", err)
		return nil, err
	}

	return networks, nil
}
