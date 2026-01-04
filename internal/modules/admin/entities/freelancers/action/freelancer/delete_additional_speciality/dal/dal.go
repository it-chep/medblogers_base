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
func (r *Repository) DeleteAdditionalSpeciality(ctx context.Context, freelancerID, specialityID int64) error {
	sql := `
		delete from freelancer_speciality_m2m where freelancer_id = $1 and speciality_id = $2
		`

	_, err := r.db.Exec(ctx, sql, freelancerID, specialityID)
	return err
}
