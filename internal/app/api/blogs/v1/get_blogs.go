package v1

import (
	"context"
	"medblogers_base/internal/modules/blogs/action/get_blogs/dto"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"

	"github.com/samber/lo"
)

func (i *Implementation) GetBlogs(ctx context.Context, req *desc.GetBlogsRequest) (*desc.GetBlogsResponse, error) {
	resp, err := i.blogs.Actions.GetBlogs.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetBlogsResponse{
		Blogs: lo.Map(resp.Blogs, func(item dto.Blog, index int) *desc.GetBlogsResponse_BlogMiniatures {
			return &desc.GetBlogsResponse_BlogMiniatures{
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
