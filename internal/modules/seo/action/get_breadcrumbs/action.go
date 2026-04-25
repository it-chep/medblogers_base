package get_breadcrumbs

import (
	"context"
	"strings"

	"medblogers_base/internal/modules/seo/action/get_breadcrumbs/dal"
	"medblogers_base/internal/modules/seo/action/get_breadcrumbs/dto"
	"medblogers_base/internal/pkg/postgres"
)

// Repository ...
type Repository interface {
	GetBreadcrumbs(ctx context.Context, path string) (dto.Breadcrumbs, error)
}

// Action ...
type Action struct {
	repository Repository
}

// New ...
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		repository: dal.NewRepository(pool),
	}
}

// Do получение хлебных крошек возможно стоит сделать nested sets - https://habr.com/ru/articles/153861/ .
func (a *Action) Do(ctx context.Context, path string) (dto.Breadcrumbs, error) {
	normalizedPath := normalizePath(path)
	if normalizedPath == "" {
		return dto.Breadcrumbs{}, nil
	}

	return a.repository.GetBreadcrumbs(ctx, normalizedPath)
}

func normalizePath(path string) string {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return ""
	}

	if !strings.HasPrefix(trimmed, "/") {
		trimmed = "/" + trimmed
	}

	if len(trimmed) > 1 {
		trimmed = strings.TrimRight(trimmed, "/")
	}

	return trimmed
}
