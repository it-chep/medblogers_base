package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/action/blog/action/get_categories/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
)

// GetBlogCategories получение всех категорий статей
func (i *Implementation) GetBlogCategories(ctx context.Context, _ *desc.GetBlogCategoriesRequest) (resp *desc.GetBlogCategoriesResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/categories", func(ctx context.Context) error {
		resp = &desc.GetBlogCategoriesResponse{}

		categories, err := i.admin.Actions.BlogModule.GetCategories.Do(ctx)
		if err != nil {
			return err
		}

		resp.Categories = lo.Map(categories, func(item dto.BlogCategory, index int) *desc.GetBlogCategoriesResponse_Category {
			return &desc.GetBlogCategoriesResponse_Category{
				Id:   item.ID,
				Name: item.Name,
			}
		})

		return nil
	})
}
