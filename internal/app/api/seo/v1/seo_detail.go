package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
)

func (i *Implementation) GetDoctorSeoData(ctx context.Context, req *desc.GetDoctorSeoDataRequest) (*desc.GetDoctorSeoDataResponse, error) {
	response, err := i.doctors.Actions.GetSeoDetail.Do(ctx, req.DoctorSlug)
	if err != nil {
		return nil, err
	}

	return &desc.GetDoctorSeoDataResponse{
		Title:       response.Title,
		Description: response.Description,
	}, nil
}
