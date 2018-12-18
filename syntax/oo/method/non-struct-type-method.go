package main

import (
	"fmt"
	"math"
)

type MyFloat64 float64

// can not declare a method with a receiver whose type is defined in another package
// can not declare receiver of built-in types
func (f MyFloat64) Abs() float64 {
	ret := float64(f)
	if ret < 0 {
		return -ret
	}
	return ret
}

func main() {
	var f MyFloat64 = -math.Sqrt2
	fmt.Println(f)
	fmt.Println(f.Abs())
}
