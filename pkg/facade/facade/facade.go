package facade

import (
	"patterns/pkg/facade/types"
)

type Coffer interface {
	DoCoffee(t types.CoffeeType) (coffee types.Coffee, err error)
}
