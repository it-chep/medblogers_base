package dal

import (
	"context"
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

// ChangeDoctorVipActivity изменение активности випки
func (r *Repository) ChangeDoctorVipActivity(ctx context.Context, doctorID int64, activity bool) error {
	sql := `
		update docstar_site_doctor set is_vip = $1 where doctor_id = $2
	`

	_, err := r.db.Exec(ctx, sql, activity, doctorID)
	if err != nil {
		return err
	}
	return nil
}
