package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// CheckPathPermission проверяет есть ли у данной роли доступ к данному урлу
func (r *Repository) CheckPathPermission(ctx context.Context, roleID int64, path string) (bool, error) {
	sql := `
		select exists(select 1
					  from roles_permissions rp
							   join permissions p on rp.permission_id = p.id
					  where rp.role_id = $1
						and p.url ilike $2)
	`
	var exists bool
	err := pgxscan.Get(ctx, r.db, &exists, sql, roleID, path)
	if err != nil {
		return false, err
	}
	return exists, nil
}
