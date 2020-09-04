package facade

import (
	"patterns/pkg/cap"
	"patterns/pkg/coffee"
	"patterns/pkg/milk"
	"patterns/pkg/water"
)

const (
	Espresso = iota + 1
	Cappuccino
	Latte
)

type Coffer interface {
	DoCoffee(t CoffeeType) (coffee Coffee, err error)
}

type CoffeeType int

type Coffee struct {
	Name       string
	CoffeeType CoffeeType
	Volume     int
	IsMilk     bool
	IsSugar    bool
}

type coffeeSvc struct {
	milk   milk.MilkWork
	coffee coffee.CoffeeWork
	water  water.WaterWork
}

func (c *coffeeSvc) DoCoffee(t CoffeeType) (coffee Coffee, err error) {
	coffee = Coffee{CoffeeType: t}
	cap := cap.Cap{}
	switch t {
	case Espresso:
		coffee.Name = "Espresso"
		c.coffee.AddCoffee(&cap)
		c.water.AddWater(&cap)
	case Cappuccino:
		coffee.Name = "Cappuccino"
		c.coffee.AddCoffee(&cap)
		c.water.AddWater(&cap)
		c.milk.AddMilk(&cap)
	case Latte:
		coffee.Name = "Latte"
		c.coffee.AddCoffee(&cap)
		c.water.AddWater(&cap)
		c.milk.AddMilk(&cap)
	}
	coffee.Volume = cap.Volume
	coffee.IsMilk = cap.IsMilk
	coffee.IsSugar = cap.IsSugar
	return
}

func New(coffee coffee.CoffeeWork, milk milk.MilkWork, water water.WaterWork) (s *coffeeSvc) {
	return &coffeeSvc{
		coffee: coffee,
		milk:   milk,
		water:  water}
}
