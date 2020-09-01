package types

type Circle struct {
	Dot
	Radius int
}

func (c Circle) Accept(visitor Visitor) (s string) {
	s = visitor.VisitCircle(c)
	return
}
