package v1

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
)

// PublishBlog публикация статьи
func (i *Implementation) PublishBlog(ctx context.Context, req *desc.PublishBlogRequest) (resp *desc.PublishBlogResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/publish", func(ctx context.Context) error {
		resp = &desc.PublishBlogResponse{}

		err := i.admin.Actions.BlogModule.PublishBlog.Do(ctx, uuid.MustParse(req.GetBlogId()))
		if err != nil {
			return err
		}

		return nil
	})
}
