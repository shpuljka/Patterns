package rectangle

import (
	"patterns/pkg/ivisitor"
)

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Heigth int
	ID     int
}

func (r Rectangle) Accept(visitor ivisitor.Visitor) (s string) {
	s = visitor.VisitRectangle(r)
	return
}
