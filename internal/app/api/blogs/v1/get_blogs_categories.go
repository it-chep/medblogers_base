package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/modules/blogs/domain/category"
	desc "medblogers_base/internal/pb/medblogers_base/api/blogs/v1"
)

func (i *Implementation) GetBlogsCategories(ctx context.Context, req *desc.GetBlogsCategoriesRequest) (_ *desc.GetBlogsCategoriesResponse, _ error) {
	resp, err := i.blogs.Actions.GetBlogsCategories.Do(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetBlogsCategoriesResponse{
		Categories: lo.Map(resp.Categories, func(item *category.Category, _ int) *desc.GetBlogsCategoriesResponse_Category {
			return &desc.GetBlogsCategoriesResponse_Category{
				Id:         item.ID(),
				Name:       item.Name(),
				FontColor:  item.FontColor(),
				BgColor:    item.BgColor(),
				BlogsCount: resp.CategoryCountMap[item.ID()],
			}
		}),
		AllBlogsCount: resp.AllBlogsCount,
	}, nil
}
