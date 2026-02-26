package vip_card

import (
	"encoding/json"
	"time"
)

func (v *VipCard) Json() ([]byte, error) {
	data := map[string]interface{}{
		"doctor_id":              v.doctorID,
		"can_barter":             v.canBarter,
		"can_buy_advertising":    v.canBuyAdvertising,
		"can_sell_advertising":   v.canSellAdvertising,
		"short_message":          v.shortMessage,
		"blog_info":              v.blogInfo,
		"advertising_price_from": v.advertisingPriceFrom,
		"end_date":               v.endDate.Format(time.RFC3339),
	}

	return json.Marshal(data)
}
