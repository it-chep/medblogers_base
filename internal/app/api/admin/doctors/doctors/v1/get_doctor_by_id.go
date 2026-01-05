package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/doctors/doctors/v1"
)

func (i *Implementation) GetDoctorByID(ctx context.Context, req *desc.GetDoctorByIDRequest) (resp *desc.GetDoctorByIDResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/doctor/{id}", func(ctx context.Context) error {
		docDTO, err := i.admin.Actions.DoctorModule.DoctorAgg.GetDoctorByID.Do(ctx, req.GetDoctorId())
		if err != nil {
			return err
		}

		resp = &desc.GetDoctorByIDResponse{
			Id:   docDTO.ID,
			Name: docDTO.Name,
			Slug: docDTO.Slug,
			// todo остальные поля
			Image: docDTO.Image,
		}
		return nil
	})
}
