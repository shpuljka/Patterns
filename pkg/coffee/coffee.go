package coffee

import (
	"patterns/pkg/cap"
)

//Интерфейс взаимодействия с кофе
type CoffeeWork interface {
	//Добавить кофе
	AddCoffee(cap *cap.Cap)
}

type coffee struct {
	volume int
}

func (c *coffee) AddCoffee(cap *cap.Cap) {
	cap.Volume += c.volume
}

//Создать кофе
func NewCoffee(vol int) (s *coffee) {
	return &coffee{volume: vol}
}
