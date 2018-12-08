package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// type assertion. sounds like explicit type conversion in Java
	s := i.(string)
	fmt.Println(s)

	// use this to get 
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/* Note that if ok is not used, panic occurs

	f = i.(float64)		// panic
	fmt.Println(f)

	*/
}
