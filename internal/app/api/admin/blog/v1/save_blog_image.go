package v1

import (
	"context"
	"encoding/base64"
	"log"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/blogs/v1"

	"github.com/google/uuid"
)

func (i *Implementation) SaveBlogImage(ctx context.Context, req *desc.SaveBlogImageRequest) (resp *desc.SaveBlogImageResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/blog/{id}/save_image", func(ctx context.Context) error {
		data, err := base64.StdEncoding.DecodeString(req.GetImageData()) //base64 string
		if err != nil {
			log.Fatal("Ошибка декодирования:", err)
			return err
		}

		imageID, imageURL, err := i.admin.Actions.BlogModule.SaveBlogImage.Do(ctx, uuid.MustParse(req.GetBlogId()), data)
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
