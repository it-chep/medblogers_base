package get_all_cities

import (
	"context"
	"medblogers_base/internal/modules/freelancers/dal/city_dal"
	"medblogers_base/internal/modules/freelancers/domain/city"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Storage interface {
	GetAllCities(ctx context.Context) ([]*city.City, error)
}

// Action список городов для регистрации
type Action struct {
	storage Storage
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		storage: city_dal.NewRepository(pool),
	}
}

// Do выполнение
func (a Action) Do(ctx context.Context) ([]*city.City, error) {
	logger.Message(ctx, "[Reg] Получение городов для регистрации")
	cities, err := a.storage.GetAllCities(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка при получении городов для регистрации", err)
		return nil, err
	}

	return cities, nil
}
