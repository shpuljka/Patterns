package main

import (
	"fmt"
	"patterns/pkg/coffee"
	"patterns/pkg/facade"
	"patterns/pkg/milk"
	"patterns/pkg/model"
	"patterns/pkg/water"
)

func main() {

	svcCoffee := facade.NewFacade(coffee.NewCoffee(10), milk.NewMilk(100), water.NewWater(100))

	c, err := svcCoffee.DoCoffee(model.Espresso)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
	c, err = svcCoffee.DoCoffee(model.Cappuccino)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
	c, err = svcCoffee.DoCoffee(model.Latte)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
}
