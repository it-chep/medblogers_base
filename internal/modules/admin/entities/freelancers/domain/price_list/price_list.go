package price_list

type PriceListItem struct {
	id      int64
	name    string
	price   int64
	priceTo *int64
}

type PriceList []PriceListItem

func NewPriceListItem(id int64, name string, price int64, priceTo *int64) PriceListItem {
	return PriceListItem{id: id, name: name, price: price, priceTo: priceTo}
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

func (p PriceListItem) GetPriceTo() *int64 {
	return p.priceTo
}
