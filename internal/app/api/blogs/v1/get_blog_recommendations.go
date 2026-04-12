package v1

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blog_recommendations/dto"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"
	"strconv"

	"github.com/samber/lo"
)

func (i *Implementation) GetBlogRecommendations(ctx context.Context, req *desc.GetBlogRecommendationsRequest) (*desc.GetBlogRecommendationsResponse, error) {
	resp, err := i.blogs.Actions.GetBlogRecommendations.Do(ctx, req.GetBlogSlug())
	if err != nil {
		return nil, err
	}

	return &desc.GetBlogRecommendationsResponse{
		Blogs: lo.Map(resp.Blogs, func(item dto.Blog, _ int) *desc.BlogMiniatures {
			return &desc.BlogMiniatures{
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				PreviewText: item.GetPreviewText(),
				CreatedAt:   converter.FormatDateRussian(item.GetCreatedAt()),
				PhotoLink:   item.GetPrimaryPhotoURL(),
				ViewsCount:  strconv.FormatInt(item.GetViewsCount(), 10),
				Categories: lo.Map(item.Categories, func(item *category.Category, _ int) *desc.Category {
					return &desc.Category{
						Id:        item.ID(),
						Name:      item.Name(),
						FontColor: item.FontColor(),
						BgColor:   item.BgColor(),
					}
				}),
			}
		}),
	}, nil
}
