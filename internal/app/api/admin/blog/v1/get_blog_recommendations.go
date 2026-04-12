package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	blogDTO "medblogers_base/internal/modules/admin/entities/blog/action/get_blog_recommendations/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func (i *Implementation) GetBlogRecommendations(ctx context.Context, req *desc.GetBlogRecommendationsRequest) (resp *desc.GetBlogRecommendationsResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/recommendations", func(ctx context.Context) error {
		recommendations, err := i.admin.Actions.BlogModule.GetBlogRecommendations.Do(ctx, uuid.MustParse(req.GetBlogId()))
		if err != nil {
			return err
		}

		resp = &desc.GetBlogRecommendationsResponse{
			Recommendations: lo.Map(recommendations, func(item blogDTO.Recommendation, _ int) *desc.GetBlogRecommendationsResponse_Recommendation {
				return &desc.GetBlogRecommendationsResponse_Recommendation{
					BlogId: item.BlogID.String(),
					Title:  item.Title,
					Slug:   item.Slug,
				}
			}),
		}

		return nil
	})
}
