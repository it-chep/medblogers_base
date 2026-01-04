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

func (r *Repository) AddAdditionalSpeciality(ctx context.Context, freelancerID, specialityID int64) error {
	sql := `
		insert into freelancer_speciality_m2m (speciality_id, freelancer_id) values ($1, $2)
		`

	_, err := r.db.Exec(ctx, sql, specialityID, freelancerID)
	return err
}
