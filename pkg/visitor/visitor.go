package visitor

import (
	"fmt"
	"patterns/pkg/visitor/types"
	"strings"
)

type VisitorString struct {
	types.Visitor
}

func (v VisitorString) VisitDot(d types.Dot) string {
	return fmt.Sprint("Dot: id = ", d.ID, "; x = ", d.X, "; y = ", d.Y)
}

func (v VisitorString) VisitCircle(c types.Circle) string {
	return fmt.Sprint("Circle: id = ", c.ID, "; x = ", c.X, "; y = ", c.Y, "; radius = ", c.Radius)
}

func (v VisitorString) VisitRectangle(r types.Rectangle) string {
	return fmt.Sprint("Rectangle: id = ", r.ID, "; x = ", r.X, "; y = ", r.Y,
		"; width = ", r.Width, "; height = ", r.Heigth)
}

func (v VisitorString) Export(s ...types.Shape) string {
	var sb strings.Builder

	for _, shape := range s {
		sb.WriteString(shape.Accept(v))
		sb.WriteString("\n")
	}

	return sb.String()
}
