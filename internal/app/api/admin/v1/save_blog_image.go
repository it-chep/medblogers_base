package v1

import (
	"context"
	"github.com/google/uuid"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/v1"
	pkgctx "medblogers_base/internal/pkg/context"
)

func (i *Implementation) SaveBlogImage(ctx context.Context, req *desc.SaveBlogImageRequest) (resp *desc.SaveBlogImageResponse, _ error) {
	email := pkgctx.GetEmailFromContext(ctx)
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, email, "/api/v1/auth/blog/{id}/save_image", func(ctx context.Context) error {
		imageID, imageURL, err := i.admin.Actions.BlogModule.SaveBlogImage.Do(ctx, uuid.MustParse(req.GetBlogId()), req.GetImageData())
		if err != nil {
			return err
		}

		resp = &desc.SaveBlogImageResponse{
			Image: &desc.SaveBlogImageResponse_Image{
				ImageUrl: imageURL,
				ImageId:  imageID.String(),
			},
		}

		return nil
	})
}
