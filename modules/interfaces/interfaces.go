package interfaces

import (
	"curso_basico_go/modules/functions"
	"math"
)

type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return math.Pow(s.base, 2)
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figure2D) {
	functions.MyPrint("Area: ", f.area())
}

func MyInterfaces() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// Interfaces list
	myInterface := []interface{}{1, true, "Hi", 15.05, math.Pi}
	functions.MyPrint(myInterface...)
}
