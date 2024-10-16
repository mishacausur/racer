package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) Area() float64 {
	return float64(r.height * r.width)
}

func main() {
	figure_1 := Circle{radius: 1.0}
	fmt.Println(figure_1.Area())
	figure_2 := Rectangle{width: 57.2, height: 10.2}
	fmt.Println(figure_2.Area())
}
