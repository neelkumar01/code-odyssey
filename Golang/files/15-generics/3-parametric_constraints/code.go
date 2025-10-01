package main

import "fmt"

type product interface {
	Name() string
	Price() float64
}

type book struct {
	title string
	author string
	price float64
}
func (b book) Name()string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}
func (b book) Price() float64 {
	return b.price
}

type toy struct{
	name string
	price float64
}
func (t toy) Name() string {
	return t.name
}
func (t toy) Price() float64 {
	return t.price
}

func applyDiscount[P product](p product, discountPercent float64) float64 {
	original := p.Price()
	discount := original * (discountPercent/100)
	return  original - discount
}

func main() {
	b := book{
		title: "The Secret",
		author: "Rhonda Byrne",
		price: 200.0,
	}

	bookDiscounted := applyDiscount[book](b, 30)

	fmt.Printf("Price of %s now: %.2f", b.Name(), bookDiscounted)
}