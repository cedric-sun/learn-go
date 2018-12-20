package main

import "fmt"

// basic function
func add(x int, y int) int {
	return x+y
}

// combined parameters list
func subtract(x,y int) int {
	return x-y
}

// multiple return
func swap(x,y int) (int,int) {
	return y,x
}

// named return values
func timesAndDiv(x,y int) (product int, quotient float64){
	product = x*y
	quotient = float64(x)/float64(y)
	// just `return` is enough for named return value
	return
}

func main() {
	fmt.Println(add(3,5))
	fmt.Println(subtract(10,7))
	fmt.Println(swap(10,20))
	fmt.Println(timesAndDiv(100,32))
}
