package v1

import (
	"context"
	"encoding/base64"
	"log"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/v1"
)

func (i *Implementation) SaveDoctorPhoto(ctx context.Context, req *desc.SaveDoctorPhotoRequest) (resp *desc.SaveDoctorPhotoResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}/save_photo", func(ctx context.Context) error {
		data, err := base64.StdEncoding.DecodeString(req.GetImageData()) //base64 string
		if err != nil {
			log.Fatal("Ошибка декодирования:", err)
			return err
		}

		imageURL, err := i.admin.Actions.DoctorModule.DoctorAgg.SaveDoctorPhoto.Do(ctx, req.GetDoctorId(), data)
		if err != nil {
			return err
		}

		resp = &desc.SaveDoctorPhotoResponse{
			ImageUrl: imageURL,
		}

		return nil
	})
}
