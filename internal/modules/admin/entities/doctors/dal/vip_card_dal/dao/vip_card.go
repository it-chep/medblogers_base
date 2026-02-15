package dao

import (
	"database/sql"
	"medblogers_base/internal/modules/admin/entities/doctors/domain/vip_card"
)

type VipMiniature struct {
	DoctorID             int64         `db:"doctor_id"`
	CanBarter            sql.NullBool  `db:"can_barter"`
	CanBuyAdvertising    sql.NullBool  `db:"can_buy_advertising"`
	CanSellAdvertising   sql.NullBool  `db:"can_sell_advertising"`
	AdvertisingPriceFrom sql.NullInt64 `db:"advertising_price_from"`
}

func (v *VipMiniature) ToDomain() *vip_card.VipCard {
	return vip_card.New(
		vip_card.WithCanBarter(v.CanBarter.Bool),
		vip_card.WithCanBuyAdvertising(v.CanBuyAdvertising.Bool),
		vip_card.WithCanSellAdvertising(v.CanSellAdvertising.Bool),
		vip_card.WithAdvertisingPriceFrom(v.AdvertisingPriceFrom.Int64),
	)
}

type VipCardDao struct {
	CanBarter          sql.NullBool `db:"can_barter"`
	CanBuyAdvertising  sql.NullBool `db:"can_buy_advertising"`
	CanSellAdvertising sql.NullBool `db:"can_buy_advertising"`

	ShortMessage         sql.NullString `db:"short_message"`
	BlogInfo             sql.NullString `db:"blog_info"`
	AdvertisingPriceFrom sql.NullInt64  `db:"advertising_price_from"`

	EndDate sql.NullTime `db:"end_date"`
}

func (v *VipCardDao) ToDomain() *vip_card.VipCard {
	return vip_card.New(
		vip_card.WithCanBarter(v.CanBarter.Bool),
		vip_card.WithCanBuyAdvertising(v.CanBuyAdvertising.Bool),
		vip_card.WithCanSellAdvertising(v.CanSellAdvertising.Bool),
		vip_card.WithShortMessage(v.ShortMessage.String),
		vip_card.WithBlogInfo(v.BlogInfo.String),
		vip_card.WithAdvertisingPriceFrom(v.AdvertisingPriceFrom.Int64),
		vip_card.WithEndDate(v.EndDate.Time),
	)
}
