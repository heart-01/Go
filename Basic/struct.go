package main

import "fmt"

type product struct {
	name  string
	price float64
	stock int
}

func (p product) clear() product {
	p.name = ""
	p.price = 0
	p.stock = 0
	return p
}

func (p product) setDiscount(discount float64) product {
	p.price -= discount
	return p
}

func main() {
	var p1 product
	p1.name = "Arduino"
	p1.price = 100.00
	p1.stock = 20

	showProduct(p1)
	updateProduct(&p1)
	showProduct(p1)

	p1 = p1.setDiscount(10.00)
	showProduct(p1)

	p1 = p1.clear()
	showProduct(p1)
}

func showProduct(p product) {
	fmt.Println(p.name, p.price, p.stock)
}

func updateProduct(p *product) {
	p.price = 200.00
	p.stock += 10
	fmt.Println(*p)
}
