package main

import "fmt"

type Vertex struct {
	X,Y float64
}

// value receiver
func (v Vertex) Scale(f float64) {
	// v here is a different instance from the argument
	// work like C++'s object-argument-pass-as-value
	// This holds true not only for receiver, but for all function argument
	v.X *= f
	v.Y *= f
}

// pointer receiver
func (v *Vertex) ScalePointer(f float64) {
	v.X *= f
	v.Y *= f
}

// function version
func ScalePointerFunction(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3,4}
	v.Scale(10)	// will not change v
	fmt.Println(v.X, v.Y)

	ptr := &v
	ptr.ScalePointer(10)
	fmt.Println(v.X, v.Y)

	v.ScalePointer(10)	// can call directly on v
	fmt.Println(v.X, v.Y)

	ScalePointerFunction(&v, 0.5)	// function version must take &v
	fmt.Println(v.X, v.Y)

	// value receiver can take pointer argument, but also will not change v
	ptr.Scale(0.5)
	fmt.Println(v.X, v.Y)
}
