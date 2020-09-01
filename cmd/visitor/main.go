package main

import (
	"fmt"
	"patterns/pkg/visitor"
	"patterns/pkg/visitor/types"
)

func main() {
	dot := types.Dot{ID: 1, X: 10, Y: 55}
	circle := types.Circle{types.Dot{ID: 2, X: 23, Y: 15}, 10}
	rectangle := types.Rectangle{ID: 3, X: 10, Y: 17, Heigth: 20, Width: 30}
	v := visitor.VisitorString{}
	fmt.Println(v.Export(dot, rectangle, circle))
}
