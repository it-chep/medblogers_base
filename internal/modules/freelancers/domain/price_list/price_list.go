package price_list

type PriceListItem struct {
	name  string
	price int64
}

type PriceList []PriceListItem

func NewPriceListItem(name string, price int64) PriceListItem {
	return PriceListItem{name, price}
}

func (p PriceListItem) GetName() string {
	return p.name
}

func (p PriceListItem) GetPrice() int64 {
	return p.price
}
