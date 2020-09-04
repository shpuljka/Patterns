package circle

import (
	"patterns/pkg/dot"
	"patterns/pkg/ivisitor"
)

type Circle struct {
	dot.Dot
	Radius int
}

func (c Circle) Accept(visitor ivisitor.Visitor) (s string) {
	s = visitor.VisitCircle(c)
	return
}
