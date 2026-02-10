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

// DeleteDoctorAdditionalSpeciality удаляет дополнителную специальность врачу
func (r *Repository) DeleteDoctorAdditionalSpeciality(ctx context.Context, doctorID, specialityID int64) error {
	sql := `
		delete from docstar_site_doctor_additional_specialties 
		       where doctor_id = $1 and speciallity_id = $2
		`

	_, err := r.db.Exec(ctx, sql, doctorID, specialityID)
	return err
}
