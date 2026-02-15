package vip_card

type VipCard struct {
	canBarter            bool
	canBuyAdvertising    bool
	canSellAdvertising   bool
	shortMessage         string
	blogInfo             string
	advertisingPriceFrom int64
}

func New(options ...Option) *VipCard {
	d := &VipCard{}
	for _, option := range options {
		option(d)
	}
	return d
}
