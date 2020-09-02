package main

import (
	"fmt"
	"patterns/pkg/facade/facade"
	"patterns/pkg/facade/types"
)

func main() {

	svcCoffee := facade.New()

	c, err := svcCoffee.DoCoffee(types.Espresso)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
	c, err = svcCoffee.DoCoffee(types.Cappuccino)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
	c, err = svcCoffee.DoCoffee(types.Latte)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(c)
}
