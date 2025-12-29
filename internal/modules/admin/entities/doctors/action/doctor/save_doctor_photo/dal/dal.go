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

func (r *Repository) SaveDoctorImage(ctx context.Context, doctorID int64, image string) error {
	sql := `
		update docstar_site_doctor 
		set s3_image = $1 
		where doctor_id = $2
	`

	_, err := r.db.Exec(ctx, sql, image, doctorID)
	return err
}
