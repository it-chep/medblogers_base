package category

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
