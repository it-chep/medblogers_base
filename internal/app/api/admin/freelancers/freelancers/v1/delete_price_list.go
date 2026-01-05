package v1

import (
	"context"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) DeletePriceList(ctx context.Context, req *desc.DeletePriceListRequest) (resp *desc.DeletePriceListResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/delete_price_list", func(ctx context.Context) error {
		return i.admin.Actions.FreelancerModule.FreelancerAgg.DeletePriceList.Do(ctx, req.GetFreelancerId(), req.GetPriceListItemId())
	})
}
