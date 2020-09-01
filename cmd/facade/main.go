package main

import (
	"fmt"
	"patterns/pkg/facade/service"
	"patterns/pkg/facade/storage"
	"patterns/pkg/facade/types"
)

func main() {
	cm, err := storage.NewStorage()
	if err != nil {
		fmt.Println("Error")
		return
	}
	var svcCoffee *service.CoffeeSvc
	svcCoffee, err = service.New(*cm)
	if err != nil {
		fmt.Println("Error")
		return
	}
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
