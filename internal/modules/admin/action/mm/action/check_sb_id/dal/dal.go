package dal

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"medblogers_base/internal/modules/admin/action/mm/action/check_sb_id/dto"
	"medblogers_base/internal/pkg/postgres"
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

// GetEmptySBIDUsers получаем пользаков без SB_ID
func (r *Repository) GetEmptySBIDUsers(ctx context.Context) (dto.GetcourseUsers, error) {
	sql := `
		select * 
			from getcourse_users 
			where sb_id is null
	`

	var users dto.GetcourseUsers
	err := pgxscan.Select(ctx, r.db, &users, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}
