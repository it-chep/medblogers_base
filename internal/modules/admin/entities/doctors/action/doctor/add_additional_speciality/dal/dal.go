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

// AddDoctorAdditionalSpeciality добавляет дополнителную специальность врачу
func (r *Repository) AddDoctorAdditionalSpeciality(ctx context.Context, doctorID, specialityID int64) error {
	sql := `
		insert into docstar_site_doctor_additional_specialties (doctor_id, speciallity_id) 
		values ($1, $2)
	`

	_, err := r.db.Exec(ctx, sql, doctorID, specialityID)
	return err
}
