package facade

import (
	"patterns/pkg/model"
)

type milkWork interface {
	AddMilk(cap *model.Cup)
}

type coffeeWork interface {
	AddCoffee(cap *model.Cup)
}

type waterWork interface {
	AddWater(cap *model.Cup)
}

//Интерфейс взаимодействия с фасадом
type Coffer interface {
	//Сделать кофе типа t
	DoCoffee(t model.CoffeeType) (coffee model.Coffee, err error)
}

type coffeeSvc struct {
	milk   milkWork
	coffee coffeeWork
	water  waterWork
}

func (c *coffeeSvc) DoCoffee(t model.CoffeeType) (coffee model.Coffee, err error) {
	coffee = model.Coffee{CoffeeType: t}
	cap := model.Cup{}
	switch t {
	case model.Espresso:
		coffee.Name = "Espresso"
		c.coffee.AddCoffee(&cap)
		c.water.AddWater(&cap)
	case model.Cappuccino:
		coffee.Name = "Cappuccino"
		c.coffee.AddCoffee(&cap)
		c.water.AddWater(&cap)
		c.milk.AddMilk(&cap)
	case model.Latte:
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

//Создать фасад
func NewFacade(coffee coffeeWork, milk milkWork, water waterWork) (s *coffeeSvc) {
	return &coffeeSvc{
		coffee: coffee,
		milk:   milk,
		water:  water,
	}
}
