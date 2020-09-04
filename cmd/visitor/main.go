package main

import (
	"fmt"
	"patterns/pkg/circle"
	"patterns/pkg/dot"
	"patterns/pkg/rectangle"
	"patterns/pkg/visitor"
)

func main() {
	dots := dot.Dot{ID: 1, X: 10, Y: 55}
	circle := circle.Circle{dot.Dot{ID: 2, X: 23, Y: 15}, 10}
	rectangle := rectangle.Rectangle{ID: 3, X: 10, Y: 17, Heigth: 20, Width: 30}
	v := visitor.VisitorString{}
	fmt.Println(v.Export(dots, rectangle, circle))
}
