package dal

import (
	"context"
	"medblogers_base/internal/modules/admin/entities/doctors/action/doctor/change_vip_info/dto"
	"medblogers_base/internal/pkg/postgres"
)

type Repository struct {
	db postgres.PoolWrapper
}

// NewRepository создает новый репозиторий по работе с докторами
func NewRepository(db postgres.PoolWrapper) *Repository {
	return &Repository{
		db: db,
	}
}

// UPSERTVipInfo обновление випки или создание
func (r *Repository) UPSERTVipInfo(ctx context.Context, doctorID int64, updateVipRequest dto.UpdateRequest) error {
	sql := `
		insert into vip_card (
			doctor_id,
			can_barter, 
			can_buy_advertising, 
			can_sell_advertising, 
			short_message, 
			advertising_price_from, 
			blog_info, 
			end_date
		) values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (doctor_id) do update set
			can_barter = EXCLUDED.can_barter,
			can_buy_advertising = EXCLUDED.can_buy_advertising,
			can_sell_advertising = EXCLUDED.can_sell_advertising,
			short_message = EXCLUDED.short_message,
			advertising_price_from = EXCLUDED.advertising_price_from,
			blog_info = EXCLUDED.blog_info,
			end_date = EXCLUDED.end_date
	`

	args := []interface{}{
		doctorID,
		updateVipRequest.CanBarter,
		updateVipRequest.CanBuyAdvertising,
		updateVipRequest.CanSellAdvertising,
		updateVipRequest.ShortMessage,
		updateVipRequest.AdvertisingPriceFrom,
		updateVipRequest.BlogInfo,
		updateVipRequest.EndDate,
	}

	_, err := r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
