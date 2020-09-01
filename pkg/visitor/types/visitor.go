package types

type Visitor interface {
	VisitDot(dot Dot) string
	VisitCircle(circle Circle) string
	VisitRectangle(rectangle Rectangle) string
}

type Shape interface {
	Accept(visitor Visitor) (s string)
}
