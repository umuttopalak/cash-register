package main

import (
	"cashregist/calculate"
	"fmt"
)

func main() {
	item1 := calculate.Item{Name: "Elma", Price: 10, Discount: 20}
	item2 := calculate.Item{Name: "Armut", Price: 20, Discount: 30}
	item3 := calculate.Item{Name: "Şeftali", Price: 30, Discount: 0}

	items := []calculate.Item{item1, item2, item3}

	for _, i := range items {
		fmt.Println(i.Description())
	}

	fmt.Printf("Toplam Fiyat: %2.f TL\n", calculate.TotalPrice(items))
}
