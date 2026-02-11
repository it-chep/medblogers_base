package v1

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_top_blogs/dto"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"

	"github.com/samber/lo"
)

func (i *Implementation) GetTopBlogs(ctx context.Context, req *desc.GetTopBlogsRequest) (*desc.GetTopBlogsResponse, error) {
	resp, err := i.blogs.Actions.GetTopBlogs.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetTopBlogsResponse{
		Blogs: lo.Map(resp.Blogs, func(item dto.Blog, index int) *desc.BlogMiniatures {
			return &desc.BlogMiniatures{
				Title:       item.GetTitle(),
				Slug:        item.GetSlug(),
				PreviewText: item.GetPreviewText(),
				CreatedAt:   converter.FormatDateRussian(item.GetCreatedAt()),
				PhotoLink:   item.GetPrimaryPhotoURL(),

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
