package facade

import (
	"patterns/pkg/facade/types"
)

type localStorage interface {
	addMilk(coffee *types.Coffee, vol int)
	addSugar(coffee *types.Coffee)
	addWater(coffee *types.Coffee)
	addCoffee(coffee *types.Coffee)
}

type Coffer interface {
	DoCoffee(t types.CoffeeType) (coffee types.Coffee, err error)
}

type coffeeSvc struct {
}

func (c *coffeeSvc) DoCoffee(t types.CoffeeType) (coffee types.Coffee, err error) {
	coffee = types.Coffee{CoffeeType: t}
	switch t {
	case types.Espresso:
		coffee.Name = "Espresso"
		c.addCoffee(&coffee)
		c.addWater(&coffee)
	case types.Cappuccino:
		coffee.Name = "Cappuccino"
		c.addCoffee(&coffee)
		c.addWater(&coffee)
		c.addMilk(&coffee, 100)
	case types.Latte:
		coffee.Name = "Latte"
		c.addCoffee(&coffee)
		c.addWater(&coffee)
		c.addMilk(&coffee, 200)
		c.addSugar(&coffee)
	}
	return
}

func (c *coffeeSvc) addMilk(coffee *types.Coffee, vol int) {
	coffee.IsMilk = true
	coffee.Volume += vol
}

func (c *coffeeSvc) addSugar(coffee *types.Coffee) {
	coffee.IsSugar = true
}

func (c *coffeeSvc) addWater(coffee *types.Coffee) {
	coffee.Volume += 100
}

func (c *coffeeSvc) addCoffee(coffee *types.Coffee) {
	coffee.Volume += 10
}

func New() (s *coffeeSvc) {
	return &coffeeSvc{}
}
