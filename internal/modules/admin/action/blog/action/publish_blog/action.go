package publish_blog

import (
	"context"
	"medblogers_base/internal/modules/admin/action/blog/action/publish_blog/dal"
	"medblogers_base/internal/modules/admin/action/blog/action/publish_blog/dto"
	"medblogers_base/internal/modules/admin/action/blog/action/publish_blog/rules"
	"medblogers_base/internal/pkg/postgres"
	"medblogers_base/internal/pkg/spec"

	"github.com/google/uuid"
)

// Action .
type Action struct {
	dal *dal.Repository
}

// New .
func New(pool postgres.PoolWrapper) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

// Do публикация статьи
func (a *Action) Do(ctx context.Context, blogID uuid.UUID) error {
	blog, err := a.dal.GetBlogByID(ctx, blogID)
	if err != nil {
		return err
	}

	validationErrors := spec.NewIndependentSpecification[*dto.Blog]().
		And(rules.RuleFieldsAvailableToPublish()).
		And(rules.RuleNotPublished()).Validate(ctx, &blog)

	if len(validationErrors) > 0 {
		return validationErrors[0]
	}

	return a.dal.PublishBlog(ctx, blogID)
}
