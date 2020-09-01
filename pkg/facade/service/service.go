package service

import (
	"patterns/pkg/facade/storage"
	"patterns/pkg/facade/types"
)

type CoffeeSvc struct {
	db storage.CoffeeMachine
}

func (c *CoffeeSvc) DoCoffee(t types.CoffeeType) (coffee types.Coffee, err error) {
	coffee = types.Coffee{CoffeeType: t}
	switch t {
	case types.Espresso:
		coffee.Name = "Espresso"
		c.db.AddCoffee(&coffee)
		c.db.AddWater(&coffee)
	case types.Cappuccino:
		coffee.Name = "Cappuccino"
		c.db.AddCoffee(&coffee)
		c.db.AddWater(&coffee)
		c.db.AddMilk(&coffee, 100)
	case types.Latte:
		coffee.Name = "Latte"
		c.db.AddCoffee(&coffee)
		c.db.AddWater(&coffee)
		c.db.AddMilk(&coffee, 200)
		c.db.AddSugar(&coffee)
	}
	return
}

func New(db storage.CoffeeMachine) (s *CoffeeSvc, err error) {
	return &CoffeeSvc{
		db: db,
	}, nil
}
