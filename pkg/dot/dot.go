package dot

import "patterns/pkg/ivisitor"

type Dot struct {
	x  int
	y  int
	id int
}

func (d *Dot) X() int {
	return d.x
}

func (d *Dot) Y() int {
	return d.y
}

func (d *Dot) ID() int {
	return d.id
}

func (d *Dot) Accept(visitor ivisitor.Visitor) (s string) {
	s = visitor.VisitDot(d)
	return
}

func New(x int, y int, id int) (s *Dot) {
	return &Dot{
		x:  x,
		y:  y,
		id: id,
	}
}
