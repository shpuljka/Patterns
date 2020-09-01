package types

type Dot struct {
	X  int
	Y  int
	ID int
}

func (d Dot) Accept(visitor Visitor) (s string) {
	s = visitor.VisitDot(d)
	return
}
