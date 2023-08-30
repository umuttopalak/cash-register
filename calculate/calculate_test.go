package calculate_test

import (
	"cashregist/calculate"
	"fmt"
	"testing"
)

func TestCalculatePriceWithNoDiscount(t *testing.T) {

	item := calculate.Item{Name: "Armut", Price: 10, Discount: 0}

	want := item.Price - (item.Price*item.Discount)/100
	got := calculate.CalculatePrice(item)

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestCalculatePriceWithDiscount(t *testing.T) {

	item := calculate.Item{Name: "Elma", Price: 20, Discount: 20}

	want := item.Price - (item.Price*item.Discount)/100
	got := calculate.CalculatePrice(item)

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestTotalPrice(t *testing.T) {

	items := []calculate.Item{{Name: "x", Price: 100, Discount: 0}, {Name: "x", Price: 80, Discount: 10}, {"x", 20, 0}}

	var want float64
	for _, i := range items {
		want += calculate.CalculatePrice(i)
	}

	got := calculate.TotalPrice(items)

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}

}

func TestDescriptionWithDiscount(t *testing.T) {

	items := []calculate.Item{{Name: "Elma", Price: 100, Discount: 10}, {Name: "Armut", Price: 0, Discount: 50}, {"Patates", 20, 100}}

	var want string
	var got string

	for _, i := range items {
		want = fmt.Sprintf("%s - %2.f TL (%2.f%% İndirimli %2.f TL)", i.Name, i.Price, i.Discount, calculate.CalculatePrice(i))
		got = i.Description()

		if want != got {
			t.Errorf("Want: %s, Got: %s", want, got)
		}
	}
}

func TestDescriptionWithNoDiscount(t *testing.T) {

	items := []calculate.Item{{Name: "Elma", Price: 100, Discount: 0}, {Name: "Armut", Price: 0, Discount: 0}, {"Patates", 20, 0}}

	var want string
	var got string

	for _, i := range items {
		want = fmt.Sprintf("%s - %2.f TL", i.Name, i.Price)
		got = i.Description()

		if want != got {
			t.Errorf("Want: %s, Got: %s", want, got)
		}
	}
}
