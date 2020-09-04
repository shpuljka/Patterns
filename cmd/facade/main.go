package main

import (
	"fmt"
	"patterns/pkg/coffee"
	"patterns/pkg/facade"
	"patterns/pkg/milk"
	"patterns/pkg/water"
)

func main() {

	svcCoffee := facade.New(coffee.New(10), milk.New(100), water.New(100))

	c, err := svcCoffee.DoCoffee(facade.Espresso)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
	c, err = svcCoffee.DoCoffee(facade.Cappuccino)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
	c, err = svcCoffee.DoCoffee(facade.Latte)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
}
