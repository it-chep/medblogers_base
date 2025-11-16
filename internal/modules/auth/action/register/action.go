package register

import (
	"context"
	"medblogers_base/internal/modules/auth/domain/user"
)

type Action struct {
}

func New() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context, email string, password string) (user.User, error) {
	return user.User{}, nil
}
