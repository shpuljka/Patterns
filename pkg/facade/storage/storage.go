package storage

import "patterns/pkg/facade/types"

type CoffeeMachine struct {
}

func (c *CoffeeMachine) AddMilk(coffee *types.Coffee, vol int) {
	coffee.IsMilk = true
	coffee.Volume += vol
}

func (c *CoffeeMachine) AddSugar(coffee *types.Coffee) {
	coffee.IsSugar = true
}

func (c *CoffeeMachine) AddWater(coffee *types.Coffee) {
	coffee.Volume += 100
}

func (c *CoffeeMachine) AddCoffee(coffee *types.Coffee) {
	coffee.Volume += 10
}

func NewStorage() (c *CoffeeMachine, err error) {
	c = &CoffeeMachine{}
	return c, nil
}
