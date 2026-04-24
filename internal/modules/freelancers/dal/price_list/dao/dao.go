package dao

import (
	"database/sql"
	"medblogers_base/internal/modules/freelancers/domain/price_list"
)

type PriceListItem struct {
	Name    string        `db:"name"`
	Price   int64         `db:"price"`
	PriceTo sql.NullInt64 `db:"price_to"`
}

func (p PriceListItem) ToDomain() price_list.PriceListItem {
	var priceTo *int64
	if p.PriceTo.Valid {
		value := p.PriceTo.Int64
		priceTo = &value
	}

	return price_list.NewPriceListItem(p.Name, p.Price, priceTo)
}

type PriceList []PriceListItem

func (p PriceList) ToDomain() price_list.PriceList {
	domain := make(price_list.PriceList, 0, len(p))
	for _, priceListItem := range p {
		domain = append(domain, priceListItem.ToDomain())
	}

	return domain
}
