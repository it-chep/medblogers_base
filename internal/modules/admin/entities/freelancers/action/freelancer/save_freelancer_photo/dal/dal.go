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

func (r *Repository) SaveFreelancerImage(ctx context.Context, freelancerID int64, image string) error {
	sql := `
		update freelancer 
		set s3_image = $1 
		where id = $2
	`

	_, err := r.db.Exec(ctx, sql, image, freelancerID)
	return err
}
