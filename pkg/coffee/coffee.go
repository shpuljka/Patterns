package coffee

import (
	"patterns/pkg/model"
)

//Интерфейс взаимодействия с кофе
type CoffeeWork interface {
	//Добавить кофе
	AddCoffee(cap *model.Cup)
}

type coffee struct {
	volume int
}

func (c *coffee) AddCoffee(cap *model.Cup) {
	cap.Volume += c.volume
}

//Создать кофе
func NewCoffee(vol int) (s *coffee) {
	return &coffee{volume: vol}
}
