package model

const (
	Espresso = iota + 1
	Cappuccino
	Latte
)

//Тип кофе
type CoffeeType int

//Структура кофе
type Coffee struct {
	Name       string
	CoffeeType CoffeeType
	Volume     int
	IsMilk     bool
	IsSugar    bool
}
