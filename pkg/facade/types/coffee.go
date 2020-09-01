package types

type CoffeeType int

const (
	Espresso = iota + 1
	Cappuccino
	Latte
)

type Coffee struct {
	Name       string
	CoffeeType CoffeeType
	Volume     int
	IsMilk     bool
	IsSugar    bool
}
