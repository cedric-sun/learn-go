package main

import "fmt"

type Vertex struct {
	X,Y float64
}

// type switch
// note that v is treated as the corresponding type in each type case
func do(i interface{}) {
	switch v:= i.(type) {
	case int:
		fmt.Printf("int: twice %v is %v\n",v,v<<1)
	case string:
		fmt.Printf("string: %q is %v bytes long\n",v, len(v))
	case Vertex:
		fmt.Printf("Vertex: X: %f, Y: %f\n",v.X, v.Y)
	default:
		// v just has the same type as i here
		fmt.Printf("I don't know this type!\n")
	}
}

func main() {
	do(21)
	do("hello")
	do(Vertex{123,354})
	do(true)
}
