package v1

import (
	"context"

	"medblogers_base/internal/app/interceptor"
	desc "medblogers_base/internal/pb/medblogers_base/api/admin/promo_offers/dictionary/v1"
)

func (i *Implementation) GetCooperationTypes(ctx context.Context, _ *desc.GetCooperationTypesRequest) (resp *desc.GetCooperationTypesResponse, _ error) {
	executor := interceptor.ExecuteWithPermissions(i.auth.Actions.CheckPermissions)

	return resp, executor(ctx, "/api/v1/admin/promo_offers/cooperation_types", func(ctx context.Context) error {
		items, err := i.admin.Actions.PromoOffers.DictionaryAgg.GetCooperationTypes.Do(ctx)
		if err != nil {
			return err
		}

		resp = newGetCooperationTypesResponse(items)
		return nil
	})
}
