package dal

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"medblogers_base/internal/modules/admin/action/mm/action/get_mm_list/dto"
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

// GetMMList получение списка мастермайндов
func (r *Repository) GetMMList(ctx context.Context) ([]dto.MM, error) {
	sql := `
		select id, mm_datetime, name, state, mm_link, created_at, is_active 
		from mm 
		order by created_at desc
		`

	var mms []dto.MM
	err := pgxscan.Select(ctx, r.db, &mms, sql)
	return mms, err
}
