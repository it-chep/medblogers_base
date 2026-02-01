package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/mm/action/create_mm/dto"
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

func (r *Repository) CreateMM(ctx context.Context, req dto.CreateMMRequest) error {
	sql := `insert into mm (mm_datetime, name, mm_link, is_active, state) values ($1, $2, $3, $4, 1)`

	args := []interface{}{
		req.MMDatetime,
		req.Name,
		req.MMLink,
		true,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	return err
}
