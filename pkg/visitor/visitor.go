package visitor

import (
	"fmt"
	"patterns/pkg/ivisitor"
	"strings"
)

type VisitorString struct {
	ivisitor.Visitor
}

func (v VisitorString) VisitDot(d ivisitor.IDot) string {
	return fmt.Sprint("Dot: id = ", d.ID(), "; x = ", d.X(), "; y = ", d.Y())
}

func (v VisitorString) VisitCircle(c ivisitor.ICircle) string {
	return fmt.Sprint("Circle: id = ", c.ID(), "; x = ", c.X(), "; y = ", c.Y(), "; radius = ", c.Radius())
}

func (v VisitorString) VisitRectangle(r ivisitor.IRectangle) string {
	return fmt.Sprint("Rectangle: id = ", r.ID(), "; x = ", r.X(), "; y = ", r.Y(),
		"; width = ", r.Width(), "; height = ", r.Heigth())
}

func (v VisitorString) Export(s ...ivisitor.Shape) string {
	var sb strings.Builder

	for _, shape := range s {
		sb.WriteString(shape.Accept(v))
		sb.WriteString("\n")
	}

	return sb.String()
}
