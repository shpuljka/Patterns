package circle

import (
	"patterns/pkg/dot"
	"patterns/pkg/ivisitor"
)

type Circle struct {
	dot.Dot
	radius int
}

func (c *Circle) Radius() int {
	return c.radius
}

func (c *Circle) Accept(visitor ivisitor.Visitor) (s string) {
	s = visitor.VisitCircle(c)
	return
}

func New(dot dot.Dot, radius int) (s *Circle) {
	return &Circle{
		Dot:    dot,
		radius: radius,
	}
}
