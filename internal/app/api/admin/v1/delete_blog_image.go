package v1

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
)

func (i *Implementation) DeleteBlogImage(ctx context.Context, req *desc.DeleteBlogImageRequest) (resp *desc.DeleteBlogImageResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/auth/blog/{id}/delete_image/{id}", func(ctx context.Context) error {
		resp = &desc.DeleteBlogImageResponse{}

		blogID := uuid.MustParse(req.GetBlogId())
		imageID := uuid.MustParse(req.GetImageId())

		err := i.admin.Actions.BlogModule.DeleteBlogImage.Do(ctx, blogID, imageID)
		if err != nil {
			return err
		}

		return nil
	})
}
