package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/update/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) UpdateFreelancer(ctx context.Context, req *desc.UpdateFreelancerRequest) (resp *desc.UpdateFreelancerResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/update", func(ctx context.Context) error {
		updateReq := getUpdateFreelancerReq(req)
		err = i.admin.Actions.FreelancerModule.FreelancerAgg.UpdateFreelancer.Do(ctx, req.GetFreelancerId(), updateReq)
		if err != nil {
			return err
		}
		return nil
	})
}

func getUpdateFreelancerReq(req *desc.UpdateFreelancerRequest) dto.UpdateRequest {
	return dto.UpdateRequest{
		Name:                 req.GetName(),
		Slug:                 req.GetSlug(),
		PortfolioLink:        req.GetPortfolioLink(),
		TgURL:                req.GetTgUrl(),
		MainCityID:           req.GetMainCityId(),
		MainSpecialityID:     req.GetMainSpecialityId(),
		AgencyRepresentative: req.GetAgencyRepresentative(),
		DateStarted:          req.GetDateStarted(),
		CooperationTypeID:    req.GetCooperationTypeId(),
		PriceCategory:        req.GetPriceCategory(),
	}
}
