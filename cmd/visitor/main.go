package main

import (
	"fmt"
	"patterns/pkg/circle"
	"patterns/pkg/dot"
	"patterns/pkg/rectangle"
	"patterns/pkg/visitor"
)

func main() {
	dots := dot.New(10, 55, 1)
	d := dot.New(23, 15, 2)
	circle := circle.New(*d, 10)
	rectangle := rectangle.New(10, 17, 3, 30, 20)
	v := visitor.VisitorString{}
	fmt.Println(v.Export(dots, rectangle, circle))
}
