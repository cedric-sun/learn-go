package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (tPtr *T) M()  {
	if (tPtr == nil) {
		fmt.Println("tPtr is nil")
	}
}

func main() {
	var t *T
	var i I = t
	i.M()
	// note that i itself is NOT nil
	if i != nil {
		fmt.Println("i itself is NOT nil")
	}
}
