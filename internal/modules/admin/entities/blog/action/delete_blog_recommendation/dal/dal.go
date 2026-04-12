package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type Repository struct {
	db postgres.PoolWrapper
}

func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) DeleteRecommendation(ctx context.Context, blogID, recommendationBlogID uuid.UUID) error {
	sql := `
		delete from blogs_recommendations
		where blog_id = $1
		  and recommended_blog_id = $2
	`

	_, err := r.db.Exec(ctx, sql, blogID.String(), recommendationBlogID.String())
	return err
}
