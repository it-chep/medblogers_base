package category

import "github.com/samber/lo"

func (c *Category) ID() int64 {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) FontColor() string {
	return c.fontColor
}

func (c *Category) BgColor() string {
	return c.bgColor
}

func (c Categories) IDs() []int64 {
	return lo.Map(c, func(item *Category, _ int) int64 {
		return item.ID()
	})
}
