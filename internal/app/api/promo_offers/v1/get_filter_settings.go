package v1

import (
	"context"

	desc "medblogers_base/internal/pb/medblogers_base/api/promo_offers/v1"
)

func (i *Implementation) GetFilterSettings(ctx context.Context, _ *desc.GetFilterSettingsRequest) (*desc.GetFilterSettingsResponse, error) {
	resp, err := i.promoOffers.Actions.GetSettings.Do(ctx)
	if err != nil {
		return nil, err
	}

	return newGetFilterSettingsResponse(resp), nil
}
