package dot

import "patterns/pkg/ivisitor"

type Dot struct {
	X  int
	Y  int
	ID int
}

func (d Dot) Accept(visitor ivisitor.Visitor) (s string) {
	s = visitor.VisitDot(d)
	return
}
