package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"

	"github.com/google/uuid"
)

// AddBlogCategory добавить статье категорию
func (i *Implementation) AddBlogCategory(ctx context.Context, req *desc.AddBlogCategoryRequest) (resp *desc.AddBlogCategoryResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/add_category", func(ctx context.Context) error {
		resp = &desc.AddBlogCategoryResponse{}

		err := i.admin.Actions.BlogModule.AddBlogCategory.Do(ctx, uuid.MustParse(req.GetBlogId()), req.GetCategoryId())
		if err != nil {
			return err
		}

		return nil
	})
}
