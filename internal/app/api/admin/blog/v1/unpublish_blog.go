package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"

	"github.com/google/uuid"
)

// UnPublishBlog Снятие статьи с публикации
func (i *Implementation) UnPublishBlog(ctx context.Context, req *desc.UnPublishBlogRequest) (resp *desc.UnPublishBlogResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/unpublish", func(ctx context.Context) error {
		resp = &desc.UnPublishBlogResponse{}

		err := i.admin.Actions.BlogModule.UnPublishBlog.Do(ctx, uuid.MustParse(req.GetBlogId()))
		if err != nil {
			return err
		}

		return nil
	})
}
