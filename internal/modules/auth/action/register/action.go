package register

import (
	"context"
	"medblogers_base/internal/modules/auth/dal/users"
	"medblogers_base/internal/pkg/postgres"

	"github.com/pkg/errors"
)

type Repository interface {
	IsEmailExists(ctx context.Context, email string) (bool, error)
	SavePass(ctx context.Context, email, password string) error
}

type Action struct {
	repo Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repo: users.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, email string, password string) error {
	exists, err := a.repo.IsEmailExists(ctx, email)
	if err != nil {
		return errors.New("Ошибка получения пользователя")
	}
	if !exists {
		return errors.New("Такого email не существует")
	}

	if err := a.repo.SavePass(ctx, email, password); err != nil {
		return errors.New("Ошибка сохранения пароля")
	}

	return nil
}
