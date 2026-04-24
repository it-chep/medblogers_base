package v1

import (
	"context"
	"github.com/samber/lo"
	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/freelancers/freelancer/v1"
)

func (i *Implementation) AddPriceList(ctx context.Context, req *desc.AddPriceListRequest) (resp *desc.AddPriceListResponse, err error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions) // todo лог действия

	return resp, executor(ctx, "/api/v1/admin/freelancer/{id}/add_price_list", func(ctx context.Context) error {
		amountTO := lo.Ternary(req.GetAmountTo() != 0, lo.ToPtr(req.GetAmountTo()), nil)
		return i.admin.Actions.FreelancerModule.FreelancerAgg.AddPriceList.Do(ctx, req.GetFreelancerId(), req.GetName(), req.GetAmount(), amountTO)
	})
}
