package main

import (
	"fmt"
	"time"
)

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	p := &Parcel{}
	p.Pdt = pdt
	p.ShippedTime = time.Now()
	return p
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {
	c := &Courier{"taejin"}
	pdt := &Product{"apple", 1000, 1}
	p := c.SendProduct(pdt)
	fmt.Println("Product:", p.Delivered())
	fmt.Println("ShippedTime:", p.ShippedTime)
	fmt.Println("DeliveredTime:", p.DeliveredTime)
}
