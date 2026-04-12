package dal

import (
	"context"
	"medblogers_base/internal/pkg/postgres"

	"github.com/georgysavva/scany/pgxscan"
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

func (r *Repository) GetRecommendationsCount(ctx context.Context, blogID uuid.UUID) (int64, error) {
	sql := `select count(*) from blogs_recommendations where blog_id = $1`

	var count int64
	err := pgxscan.Get(ctx, r.db, &count, sql, blogID.String())
	return count, err
}

func (r *Repository) AddRecommendation(ctx context.Context, blogID, recommendationBlogID uuid.UUID) error {
	sql := `
		insert into blogs_recommendations (blog_id, recommended_blog_id)
		values ($1, $2)
	`

	_, err := r.db.Exec(ctx, sql, blogID.String(), recommendationBlogID.String())
	return err
}
