package v1

import (
	"context"
	"encoding/base64"
	"log"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) SaveFreelancerPhoto(ctx context.Context, req *desc.SaveFreelancerPhotoRequest) (resp *desc.SaveFreelancerPhotoResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/save_photo", func(ctx context.Context) error {
		data, err := base64.StdEncoding.DecodeString(req.GetImageData()) //base64 string
		if err != nil {
			log.Fatal("Ошибка декодирования:", err)
			return err
		}

		imageURL, err := i.admin.Actions.FreelancerModule.FreelancerAgg.SaveFreelancerPhoto.Do(ctx, req.GetFreelancerId(), data)
		if err != nil {
			return err
		}

		resp = &desc.SaveFreelancerPhotoResponse{
			Image: imageURL,
		}

		return nil
	})
}
