package formatters

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"strings"
)

func HumanPrice(price int64) string {
	return strings.Replace(humanize.Comma(price), ",", " ", -1)
}

func HumanPriceFromWithAgreement(price int64) string {
	if price <= 0 {
		return "цена по договоренности"
	}

	priceStr := strings.Replace(humanize.Comma(price), ",", " ", -1)

	return fmt.Sprintf("от %s ₽", priceStr)
}

func FormatPriceRange(price int64, priceTo *int64) string {
	if price <= 0 {
		return "по договоренности"
	}

	if priceTo == nil || *priceTo <= 0 {
		return fmt.Sprintf("от %s ₽", HumanPrice(price))
	}

	return fmt.Sprintf("%s - %s", HumanPrice(price), HumanPrice(*priceTo))
}
