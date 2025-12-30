package price_list

type PriceListItem struct {
	id    int64
	name  string
	price int64
}

type PriceList []PriceListItem

func NewPriceListItem(id int64, name string, price int64) PriceListItem {
	return PriceListItem{id: id, name: name, price: price}
}

func (p PriceListItem) GetName() string {
	return p.name
}

func (p PriceListItem) GetPrice() int64 {
	return p.price
}

func (p PriceListItem) GetID() int64 {
	return p.id
}
