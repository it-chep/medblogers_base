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

// AddDoctorAdditionalCity добавляет дополнительный город врачу1
func (r *Repository) AddDoctorAdditionalCity(ctx context.Context, doctorID, cityID int64) error {
	sql := `
		insert into docstar_site_doctor_additional_cities (doctor_id, city_id) 
		values ($1, $2)
	`

	_, err := r.db.Exec(ctx, sql, doctorID, cityID)
	return err
}
