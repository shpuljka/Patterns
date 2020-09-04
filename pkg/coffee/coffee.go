package coffee

import (
	"patterns/pkg/cap"
)

type CoffeeWork interface {
	AddCoffee(cap *cap.Cap)
}

type coffee struct {
	volume int
}

func (c *coffee) AddCoffee(cap *cap.Cap) {
	cap.Volume += c.volume
}

func New(vol int) (s *coffee) {
	return &coffee{volume: vol}
}
