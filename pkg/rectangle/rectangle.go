package rectangle

import (
	"patterns/pkg/ivisitor"
)

type Rectangle struct {
	x      int
	y      int
	width  int
	heigth int
	id     int
}

func (r *Rectangle) X() int {
	return r.x
}

func (r *Rectangle) Y() int {
	return r.y
}

func (r *Rectangle) Width() int {
	return r.width
}

func (r *Rectangle) Heigth() int {
	return r.heigth
}

func (r *Rectangle) ID() int {
	return r.id
}

func (r *Rectangle) Accept(visitor ivisitor.Visitor) (s string) {
	s = visitor.VisitRectangle(r)
	return
}

func New(x int, y int, id int, width int, height int) (s *Rectangle) {
	return &Rectangle{
		x:      x,
		y:      y,
		id:     id,
		width:  width,
		heigth: height,
	}
}
