package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDoctorSeoData(context.Context, *desc.GetDoctorSeoDataRequest) (*desc.GetDoctorSeoDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorSeoData not implemented")
}
