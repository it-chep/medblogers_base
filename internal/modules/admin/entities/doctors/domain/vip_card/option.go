package vip_card

import "time"

// Option .
type Option func(c *VipCard)

// WithCanBarter .
func WithCanBarter(canBarter bool) Option {
	return func(c *VipCard) {
		c.canBarter = canBarter
	}
}

// WithCanBuyAdvertising .
func WithCanBuyAdvertising(canBuyAdvertising bool) Option {
	return func(c *VipCard) {
		c.canBuyAdvertising = canBuyAdvertising
	}
}

// WithCanSellAdvertising .
func WithCanSellAdvertising(canSellAdvertising bool) Option {
	return func(c *VipCard) {
		c.canSellAdvertising = canSellAdvertising
	}
}

// WithShortMessage .
func WithShortMessage(shortMessage string) Option {
	return func(s *VipCard) {
		s.shortMessage = shortMessage
	}
}

// WithBlogInfo .
func WithBlogInfo(blogInfo string) Option {
	return func(s *VipCard) {
		s.blogInfo = blogInfo
	}
}

// WithAdvertisingPriceFrom .
func WithAdvertisingPriceFrom(advertisingPriceFrom int64) Option {
	return func(s *VipCard) {
		s.advertisingPriceFrom = advertisingPriceFrom
	}
}

func WithEndDate(endDate time.Time) Option {
	return func(s *VipCard) {
		s.endDate = endDate
	}
}
