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

func (r *Repository) AddRecommendation(ctx context.Context, freelancerId, doctorId int64) error {
	sql := `
		insert into freelancer_recommendation (freelancer_id, doctor_id) values ($1, $2)
		`

	_, err := r.db.Exec(ctx, sql, freelancerId, doctorId)
	return err
}
