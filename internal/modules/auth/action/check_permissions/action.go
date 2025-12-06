package check_permissions

import (
	"context"
	"medblogers_base/internal/modules/auth/action/check_permissions/dal"
	"medblogers_base/internal/modules/auth/dal/users"
	"medblogers_base/internal/modules/auth/domain/user"
	"medblogers_base/internal/pkg/postgres"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}

type Action struct {
	userRepository Repository
	dal            *dal.Repository
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		userRepository: users.NewRepository(pool),
		dal:            dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, email, path string) (*user.User, error) {
	usr, err := a.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`/\d+`)

	// Заменяем все числа в пути
	normalized := re.ReplaceAllString(path, `/{id}`)

	// Если URL заканчивается на число (например, "/tutors/1")
	if strings.HasSuffix(normalized, "/{id") && !strings.HasSuffix(normalized, "/{id}") {
		normalized += "}"
	}

	hasPermission, err := a.dal.CheckPathPermission(ctx, usr.GetRoleID(), normalized)
	if err != nil {
		return nil, err
	}

	if !hasPermission {
		return nil, errors.New("Недостаточно прав")
	}

	return usr, nil
}
