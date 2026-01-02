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

// DeleteDoctorAdditionalCity удаляет дополнительный город врачу1
func (r *Repository) DeleteDoctorAdditionalCity(ctx context.Context, doctorID, cityID int64) error {
	sql := `
		delete from docstar_site_doctor_additional_cities 
		       where doctor_id = $1 and city_id = $2
	`

	_, err := r.db.Exec(ctx, sql, doctorID, cityID)
	return err
}
