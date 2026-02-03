package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerByID(ctx context.Context, req *desc.GetFreelancerByIDRequest) (resp *desc.GetFreelancerByIDResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}", func(ctx context.Context) error {
		res, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetFreelancerByID.Do(ctx, req.GetFreelancerId())
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerByIDResponse{
			Id:   res.ID,
			Name: res.Name,
			Slug: res.Slug,
			//Email:, todo
			PortfolioLink: res.PortfolioLink,
			TgUrl:         res.TgURL,
			MainCity: &desc.CityItem{
				Id:   res.City.ID,
				Name: res.City.Name,
			},
			MainSpeciality: &desc.SpecialityItem{
				Id:   res.Speciality.ID,
				Name: res.Speciality.Name,
			},
			IsActive: res.IsActive,
			Image:    res.S3Image,
			CooperationType: &desc.CooperationType{
				Id: res.CooperationType,
				//Name:,
			},
			AgencyRepresentative: res.AgencyRepresentative,
			//CreatedAt: todo
			DateStarted:   res.StartWorking.Format("2006-01-02 15:04:05"),
			PriceCategory: res.PriceCategory,
		}
		return nil
	})
}
