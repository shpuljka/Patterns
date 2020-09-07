package ivisitor

type IDot interface {
	X() int
	Y() int
	ID() int
}

type ICircle interface {
	IDot
	Radius() int
}

type IRectangle interface {
	X() int
	Y() int
	Width() int
	Heigth() int
	ID() int
}

type Visitor interface {
	VisitDot(dot IDot) string
	VisitCircle(circle ICircle) string
	VisitRectangle(rectangle IRectangle) string
}

type Shape interface {
	Accept(visitor Visitor) (s string)
}
