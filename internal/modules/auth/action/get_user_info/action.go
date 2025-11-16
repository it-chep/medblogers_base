package get_user_info

import (
	"context"
	"medblogers_base/internal/modules/auth/domain/user"
)

type Action struct {
}

func New() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context, email string) (user.User, error) {
	return user.User{}, nil
}
