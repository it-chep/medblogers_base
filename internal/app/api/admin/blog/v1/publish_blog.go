package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"

	"github.com/google/uuid"
)

// PublishBlog публикация статьи
func (i *Implementation) PublishBlog(ctx context.Context, req *desc.PublishBlogRequest) (resp *desc.PublishBlogResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/publish", func(ctx context.Context) error {
		resp = &desc.PublishBlogResponse{}
		var imageID *uuid.UUID
		blogID := uuid.MustParse(req.GetBlogId())
		if len(req.GetPrimaryImageId()) != 0 {
			parsedUUID := uuid.MustParse(req.GetPrimaryImageId())
			imageID = &parsedUUID
		}

		err := i.admin.Actions.BlogModule.PublishBlog.Do(ctx, blogID, imageID)
		if err != nil {
			return err
		}

		return nil
	})
}
