package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/blogs/action/filter_blogs/dto"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
	"medblogers_base/internal/pkg/converter"
)

// FilterBlogs фильтрация статей
func (i *Implementation) FilterBlogs(ctx context.Context, req *desc.FilterBlogsRequest) (_ *desc.FilterBlogsResponse, _ error) {
	resp, err := i.blogs.Actions.FilterBlogs.Do(ctx, dto.FilterRequest{
		CategoriesIDs: req.GetCategoriesIds(),
	})
	if err != nil {
		return nil, err
	}

	return &desc.FilterBlogsResponse{
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
