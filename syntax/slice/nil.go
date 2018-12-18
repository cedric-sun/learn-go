package main

import "fmt"

func main() {
	// len(a) is 0, cap(a) is 0, a has no underlying array
	var a []int
	fmt.Println(a,len(a), cap(a))
	if a == nil {
		fmt.Println("a is nil!")
	}
}
