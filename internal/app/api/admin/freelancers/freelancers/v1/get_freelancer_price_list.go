package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	"medblogers_base/internal/modules/admin/entities/freelancers/action/freelancer/get_price_list/dto"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) GetFreelancerPriceList(ctx context.Context, req *desc.GetFreelancerPriceListRequest) (resp *desc.GetFreelancerPriceListResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/price_list", func(ctx context.Context) error {
		priceList, err := i.admin.Actions.FreelancerModule.FreelancerAgg.GetPriceList.Do(ctx, req.GetFreelancerId())
		if err != nil {
			return err
		}

		resp = &desc.GetFreelancerPriceListResponse{
			PriceList: lo.Map(priceList, func(item dto.PriceList, index int) *desc.PriceList {
				return &desc.PriceList{
					Id:     item.ID,
					Name:   item.Name,
					Amount: item.Amount,
				}
			}),
		}
		return nil
	})
}
