package add_blog_recommendation

import (
	"context"
	"github.com/pkg/errors"
	"medblogers_base/internal/modules/admin/entities/blog/action/add_blog_recommendation/dal"
	commonDal "medblogers_base/internal/modules/admin/entities/blog/dal"
	"medblogers_base/internal/pkg/postgres"

	"github.com/google/uuid"
)

type CommonDal interface {
	GetBlogByID(ctx context.Context, blogID uuid.UUID) error
}

type ActionDal interface {
	GetRecommendationsCount(ctx context.Context, blogID uuid.UUID) (int64, error)
	AddRecommendation(ctx context.Context, blogID, recommendationBlogID uuid.UUID) error
}

type Action struct {
	commonDal CommonDal
	actionDal ActionDal
}

func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		commonDal: commonDal.NewRepository(pool),
		actionDal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, blogID, recommendationBlogID uuid.UUID) error {
	if err := a.commonDal.GetBlogByID(ctx, blogID); err != nil {
		return err
	}

	if err := a.commonDal.GetBlogByID(ctx, recommendationBlogID); err != nil {
		return err
	}

	count, err := a.actionDal.GetRecommendationsCount(ctx, blogID)
	if err != nil {
		return err
	}

	if count >= 3 {
		return errors.New("Максимум 3 рекомендации")
	}

	return a.actionDal.AddRecommendation(ctx, blogID, recommendationBlogID)
}
