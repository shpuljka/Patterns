package types

const (
	Espresso = iota + 1
	Cappuccino
	Latte
)

type CoffeeType int

type Coffee struct {
	Name       string
	CoffeeType CoffeeType
	Volume     int
	IsMilk     bool
	IsSugar    bool
}
