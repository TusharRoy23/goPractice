package main

import "fmt"

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	(*b).price = b.price * 0.9 // b.price = b.price * 0.9
}

func main() {
	bookInfo := book{title: "ABCDEF", price: 8.1}
	vat := bookInfo.vat()
	fmt.Println("vat: ", vat)

	bookInfo.discount()
	fmt.Println(bookInfo)
}
