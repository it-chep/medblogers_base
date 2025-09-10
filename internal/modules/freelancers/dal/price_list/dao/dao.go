package dao

import "medblogers_base/internal/modules/freelancers/domain/price_list"

type PriceListItem struct {
	Name  string `db:"name"`
	Price int64  `db:"price"`
}

func (p PriceListItem) ToDomain() price_list.PriceListItem {
	return price_list.NewPriceListItem(p.Name, p.Price)
}

type PriceList []PriceListItem

func (p PriceList) ToDomain() price_list.PriceList {
	domain := make(price_list.PriceList, 0, len(p))
	for _, priceListItem := range p {
		domain = append(domain, priceListItem.ToDomain())
	}

	return domain
}
