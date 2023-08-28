package main

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
		return fmt.Sprintf("%s - %2.f TL (%2.f%% İndirimli %2.f TL)", i.Name, i.Price, i.Discount, calculatePrice(i))
	}

	return fmt.Sprintf("%s - %2.f TL", i.Name, i.Price)
}

func calculatePrice(item Item) float64 {

	if item.Discount != 0 {
		return (item.Price - ((item.Price * item.Discount) / 100))
	}

	return item.Price
}

func totalPrice(items []Item) float64 {

	var total float64

	for _, i := range items {
		total += calculatePrice(i)
	}

	return total
}

func main() {
	item1 := Item{Name: "Elma", Price: 10, Discount: 20}
	item2 := Item{Name: "Armut", Price: 20, Discount: 30}
	item3 := Item{Name: "Şeftali", Price: 30, Discount: 0}

	items := []Item{item1, item2, item3}

	for _, i := range items {
		fmt.Println(i.Description())
	}

	fmt.Println("Toplam Fiyat: ", totalPrice(items))
}
