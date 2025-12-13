package get_user_info

import (
	"context"
	"medblogers_base/internal/modules/auth/dal/users"
	"medblogers_base/internal/modules/auth/domain/user"
	"medblogers_base/internal/pkg/postgres"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}

type Action struct {
	repo Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repo: users.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, email string) (*user.User, error) {
	return a.repo.GetUserByEmail(ctx, email)
}
