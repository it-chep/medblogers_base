package vip_card

import "time"

type VipCard struct {
	doctorID int64

	canBarter            bool
	canBuyAdvertising    bool
	canSellAdvertising   bool
	shortMessage         string
	blogInfo             string
	advertisingPriceFrom int64

	endDate time.Time
}

func New(options ...Option) *VipCard {
	d := &VipCard{}
	for _, option := range options {
		option(d)
	}
	return d
}
