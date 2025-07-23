package get_all_specialities

import (
	"context"
	"medblogers_base/internal/modules/doctors/dal/speciality_dal"
	"medblogers_base/internal/modules/doctors/domain/speciality"
	"medblogers_base/internal/pkg/logger"
	"medblogers_base/internal/pkg/postgres"
)

type Storage interface {
	GetAllSpecialities(ctx context.Context) ([]*speciality.Speciality, error)
}

// Action список специальностей в для регистрации
type Action struct {
	storage Storage
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		storage: speciality_dal.NewRepository(pool),
	}
}

// Do выполнение
func (a Action) Do(ctx context.Context) ([]*speciality.Speciality, error) {
	logger.Message(ctx, "[Reg] Получение специальностей для регистрации")
	specialities, err := a.storage.GetAllSpecialities(ctx)
	if err != nil {
		logger.Error(ctx, "Ошибка при получении специальностей для регистрации", err)
		return nil, err
	}

	return specialities, nil
}
