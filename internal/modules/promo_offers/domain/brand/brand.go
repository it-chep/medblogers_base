package brand

import (
	"github.com/samber/lo"
	"time"
)

type Brand struct {
	id                 int64
	photo              string
	title              string
	slug               string
	businessCategoryID int64
	website            string
	description        string
	about              string
	isActive           bool
	createdAt          *time.Time
}

func New(options ...Option) *Brand {
	item := &Brand{}
	for _, option := range options {
		option(item)
	}

	return item
}

type Brands []*Brand

func (b Brands) IDs() []int64 {
	return lo.Map(b, func(item *Brand, _ int) int64 {
		return item.GetID()
	})
}

func (b Brands) BusinessCategoryIDs() []int64 {
	return lo.FilterMap(b, func(item *Brand, _ int) (int64, bool) {
		if item.GetBusinessCategoryID() <= 0 {
			return 0, false
		}

		return item.GetBusinessCategoryID(), true
	})
}
