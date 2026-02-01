package vip_card

// GetCanBarter .
func (v *VipCard) GetCanBarter() bool {
	return v.canBarter
}

// GetCanBuyAdvertising .
func (v *VipCard) GetCanBuyAdvertising() bool {
	return v.canBuyAdvertising
}

// GetCanSellAdvertising .
func (v *VipCard) GetCanSellAdvertising() bool {
	return v.canSellAdvertising
}

// GetShortMessage .
func (v *VipCard) GetShortMessage() string {
	return v.shortMessage
}

// GetBlogInfo .
func (v *VipCard) GetBlogInfo() string {
	return v.blogInfo
}

// GetAdvertisingPriceFrom .
func (v *VipCard) GetAdvertisingPriceFrom() int64 {
	return v.advertisingPriceFrom
}
