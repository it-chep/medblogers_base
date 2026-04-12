package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"

	"github.com/google/uuid"
)

func (i *Implementation) AddBlogRecommendation(ctx context.Context, req *desc.AddBlogRecommendationRequest) (resp *desc.AddBlogRecommendationResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/add_recommendation", func(ctx context.Context) error {
		return i.admin.Actions.BlogModule.AddBlogRecommendation.Do(
			ctx,
			uuid.MustParse(req.GetBlogId()),
			uuid.MustParse(req.GetRecommendationBlogId()),
		)
	})
}
