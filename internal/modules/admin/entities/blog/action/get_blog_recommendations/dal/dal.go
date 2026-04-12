package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/blog/action/get_blog_recommendations/dto"
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

func (r *Repository) GetRecommendations(ctx context.Context, blogID uuid.UUID) ([]dto.Recommendation, error) {
	sql := `
		select b.id as blog_id, b.name as title, b.slug
		from blogs_recommendations br
			join blog b on b.id = br.recommended_blog_id
		where br.blog_id = $1
		order by b.ordering_number
	`

	var recommendations []dto.Recommendation
	err := pgxscan.Select(ctx, r.db, &recommendations, sql, blogID.String())
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}
