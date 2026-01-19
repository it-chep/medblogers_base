package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"

	"github.com/google/uuid"
)

// DeleteBlogCategory удалить статье категорию
func (i *Implementation) DeleteBlogCategory(ctx context.Context, req *desc.DeleteBlogCategoryRequest) (resp *desc.DeleteBlogCategoryResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/delete_category", func(ctx context.Context) error {
		resp = &desc.DeleteBlogCategoryResponse{}

		err := i.admin.Actions.BlogModule.DeleteBlogCategory.Do(ctx, uuid.MustParse(req.GetBlogId()), req.GetCategoryId())
		if err != nil {
			return err
		}

		return nil
	})
}
