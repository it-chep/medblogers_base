package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AccrueMBC(ctx context.Context, doctorID, mbcCount, accruedBy int64) error {
	sql := `
		insert into mbc_operation (id, doctor_id, mbc_count, accrued_by_id)
		VALUES ($1, $2, $3, $4)
	`

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, sql, id, doctorID, mbcCount, accruedBy)
	return err
}
