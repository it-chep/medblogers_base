package v1

import (
	"context"
	desc "medblogers_base/internal/pb/medblogers_base/api/seo/v1"
)

func (i *Implementation) GetFreelancersSeoData(ctx context.Context, req *desc.GetFreelancersSeoDataRequest) (_ *desc.GetFreelancersSeoDataResponse, _ error) {
	resp, err := i.freelancer.Actions.GetSeoDetail.Do(ctx, req.GetFreelancersSlug())
	if err != nil {
		return nil, err
	}
	return &desc.GetFreelancersSeoDataResponse{
		Title:       resp.Title,
		Description: resp.Description,
		Image:       resp.ImageURL,
	}, nil
}
