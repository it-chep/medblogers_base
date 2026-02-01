package category

// Category .
type Category struct {
	id        int64
	name      string
	fontColor string
	bgColor   string
}

type Categories []*Category

// New .
func New(id int64, name, fontColor, bgColor string) *Category {
	return &Category{
		id:        id,
		name:      name,
		fontColor: fontColor,
		bgColor:   bgColor,
	}
}
