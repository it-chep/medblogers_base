package dto

import "time"

type UpdateRequest struct {
	CanBarter            bool
	CanBuyAdvertising    bool
	CanSellAdvertising   bool
	ShortMessage         string
	BlogInfo             string
	AdvertisingPriceFrom int64

	EndDate time.Time
}
