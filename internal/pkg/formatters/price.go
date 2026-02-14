package formatters

import (
	"github.com/dustin/go-humanize"
	"strings"
)

func HumanPrice(price int64) string {
	return strings.Replace(humanize.Comma(price), ",", " ", -1)
}
