package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	ret := float64(f)
	if ret <0 {
		return -ret
	}
	return ret
}

type Vertex struct {
	X,Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f		// OK, because MyFloat implements Abser
	a = &v		// OK, because *Veretx implements Abser

	// a = v	// Error, because Vertex does not implement Abser

	fmt.Println(a.Abs())
}
