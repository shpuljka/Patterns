package ivisitor

import (
	"patterns/pkg/circle"
	"patterns/pkg/dot"
	"patterns/pkg/rectangle"
)

type Visitor interface {
	VisitDot(dot dot.Dot) string
	VisitCircle(circle circle.Circle) string
	VisitRectangle(rectangle rectangle.Rectangle) string
}

type Shape interface {
	Accept(visitor Visitor) (s string)
}
