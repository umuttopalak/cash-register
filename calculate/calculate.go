package calculate

import (
	"fmt"
)

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

type Describable interface {
	Description() string
}

func (i Item) Description() string {
	if i.Discount != 0 {
		return fmt.Sprintf("%s - %2.f TL (%2.f%% İndirimli %2.f TL)", i.Name, i.Price, i.Discount, CalculatePrice(i))
	}

	return fmt.Sprintf("%s - %2.f TL", i.Name, i.Price)
}

func CalculatePrice(item Item) float64 {

	if item.Discount != 0 || item.Discount <= 100 {
		return (item.Price - ((item.Price * item.Discount) / 100))
	}

	return item.Price
}

func TotalPrice(items []Item) float64 {

	var total float64

	for _, i := range items {
		total += CalculatePrice(i)
	}

	return total
}
