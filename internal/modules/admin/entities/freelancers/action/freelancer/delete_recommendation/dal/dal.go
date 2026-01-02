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

func (r *Repository) DeleteRecommendation(ctx context.Context, freelancerID, doctorID int64) error {
	sql := `
		delete from freelancer_recommendation where freelancer_id = $1 and doctor_id = $2
		`

	_, err := r.db.Exec(ctx, sql, freelancerID, doctorID)
	return err
}
